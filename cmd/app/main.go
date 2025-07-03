package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// データ構造
type Client struct {
	ID          string
	Secret      string
	RedirectURI string
}

type User struct {
	ID       string
	Username string
	Password string
}

type AuthCode struct {
	Code      string
	ClientID  string
	UserID    string
	CreatedAt time.Time
}

// インメモリストレージ（実際はDBを使う）
var (
	clients       = make(map[string]Client)
	users         = make(map[string]User)
	authCodes     = make(map[string]AuthCode)
	jwtSecret     = []byte("your-secret-key") // 実際は環境変数から取得
	client_domain = os.Getenv("CLIENT_DOMAIN")
)

// 初期データ設定
func init() {
	// テスト用クライアント
	fmt.Println(client_domain)
	clients["test-client"] = Client{
		ID:          "test-client",
		Secret:      "test-secret",
		RedirectURI: "http://" + client_domain + "/callback",
	}

	// テスト用ユーザー
	users["user1"] = User{
		ID:       "user1",
		Username: "1",
		Password: "2", // 実際はハッシュ化する
	}
}

// JWTトークン生成
func generateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ランダム文字列生成
func generateRandomString(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// 認可エンドポイント
func authorizeHandler(w http.ResponseWriter, r *http.Request) {
	// パラメータ取得
	clientID := r.URL.Query().Get("client_id")
	redirectURI := r.URL.Query().Get("redirect_uri")
	responseType := r.URL.Query().Get("response_type")
	state := r.URL.Query().Get("state")

	// バリデーション
	client, exists := clients[clientID]
	if !exists {
		http.Error(w, "Invalid client_id", http.StatusBadRequest)
		return
	}

	if client.RedirectURI != redirectURI {
		http.Error(w, "Invalid redirect_uri", http.StatusBadRequest)
		return
	}

	if responseType != "code" {
		http.Error(w, "Unsupported response_type", http.StatusBadRequest)
		return
	}

	// ログイン画面表示
	if r.Method == "GET" {
		html := AuthorizeView()

		t, _ := template.New("login").Parse(html)
		data := map[string]string{
			"ClientID":     clientID,
			"RedirectURI":  redirectURI,
			"ResponseType": responseType,
			"State":        state,
		}
		t.Execute(w, data)
		return
	}

	// POST: 認証処理
	if r.Method == "POST" {
		username := r.FormValue("username")
		fmt.Printf("%v \n", username)
		password := r.FormValue("password")
		fmt.Printf("%v \n", password)

		// ユーザー認証（簡単な実装）
		var userID string
		for _, user := range users {
			if user.Username == username && user.Password == password {
				userID = user.ID
				break
			}
		}

		if userID == "" {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// 認可コード生成
		code := generateRandomString(16)
		authCodes[code] = AuthCode{
			Code:      code,
			ClientID:  clientID,
			UserID:    userID,
			CreatedAt: time.Now(),
		}

		// リダイレクト
		redirectURL, _ := url.Parse(redirectURI)
		q := redirectURL.Query()
		q.Set("code", code)
		if state != "" {
			q.Set("state", state)
		}
		redirectURL.RawQuery = q.Encode()

		http.Redirect(w, r, redirectURL.String(), http.StatusFound)
	}
}

// トークンエンドポイント
func tokenHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Token request received from %s", r.RemoteAddr)

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// パラメータ取得
	grantType := r.FormValue("grant_type")
	code := r.FormValue("code")
	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")
	redirectURI := r.FormValue("redirect_uri")

	log.Printf("Token request: grant_type=%s, code=%s, client_id=%s, redirect_uri=%s", grantType, code, clientID, redirectURI)

	// バリデーション
	if grantType != "authorization_code" {
		http.Error(w, "Unsupported grant_type", http.StatusBadRequest)
		return
	}

	client, exists := clients[clientID]
	if !exists || client.Secret != clientSecret {
		http.Error(w, "Invalid client", http.StatusUnauthorized)
		return
	}

	// redirect_uriの検証も追加（セキュリティ上重要）
	if client.RedirectURI != redirectURI {
		http.Error(w, "Invalid redirect_uri", http.StatusBadRequest)
		return
	}

	authCode, exists := authCodes[code]
	if !exists {
		http.Error(w, "Invalid code", http.StatusBadRequest)
		return
	}

	// 認可コードの有効期限チェック（10分）
	if time.Since(authCode.CreatedAt) > 10*time.Minute {
		delete(authCodes, code)
		http.Error(w, "Code expired", http.StatusBadRequest)
		return
	}

	// 認可コードを削除（一回限り使用）
	delete(authCodes, code)

	// アクセストークン生成
	accessToken, err := generateJWT(authCode.UserID)
	if err != nil {
		http.Error(w, "Token generation failed", http.StatusInternalServerError)
		return
	}

	// レスポンス
	response := map[string]interface{}{
		"access_token": accessToken,
		"token_type":   "Bearer",
		"expires_in":   3600,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ユーザー情報エンドポイント
func userinfoHandler(w http.ResponseWriter, r *http.Request) {
	// Authorizationヘッダーからトークン取得
	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// JWT検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// ユーザーID取得
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}

	userID := claims["sub"].(string)
	user, exists := users[userID]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// ユーザー情報レスポンス
	response := map[string]interface{}{
		"sub":      user.ID,
		"username": user.Username,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	p := ":" + os.Getenv("GO_PORT")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Health check endpoint hit from %s", r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok", "service": "oauth-Server"}`))
	})
	http.HandleFunc("/authorize", authorizeHandler)
	http.HandleFunc("/token", tokenHandler)
	http.HandleFunc("/userinfo", userinfoHandler)

	log.Printf("OAuth Server starting on %v \n", p)
	log.Fatal(http.ListenAndServe(p, nil))
}
