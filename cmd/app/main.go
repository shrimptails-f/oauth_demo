package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// データ構造
type Client struct {
	ID     string
	Secret string
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
	clients                        = make(map[string]Client)
	users                          = make(map[string]User)
	authorize_server_domain        = os.Getenv("AUTHORIZE_SERVER_DOMAIN")
	authorize_server_docker_domain = os.Getenv("AUTHORIZE_SERVER_DOCKER_DOMAIN")
)

// 初期データ設定
func init() {
	// テスト用クライアント
	clients["test-client"] = Client{
		ID:     "test-client",
		Secret: "test-secret",
	}

	// テスト用ユーザー
	users["user1"] = User{
		ID:       "user1",
		Username: "1",
		Password: "2", // 実際はハッシュ化する
	}
}

// ログイン画面ハンドラー
func loginHandler(w http.ResponseWriter, r *http.Request) {
	const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>OAuth Client - Login</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 600px; margin: 50px auto; padding: 20px; background: linear-gradient(135deg, #10b981 0%, #059669 100%);}
        .container { background: #f5f5f5; padding: 30px; border-radius: 8px; }
        h2 { color: #333; }
        .btn { 
            background: #059669; /* 深い緑背景 */
            color: white; 
            padding: 12px 24px; 
            border: none; 
            border-radius: 4px; 
            cursor: pointer; 
            text-decoration: none; 
            display: inline-block; 
            transition: all 0.3s ease;
            box-shadow: 0 2px 8px rgba(5, 150, 105, 0.2);
        }
        
        .btn:hover { 
            background: #047857; /* ホバー時はさらに深い緑 */
            box-shadow: 0 4px 12px rgba(5, 150, 105, 0.3);
            transform: translateY(-1px);
        }
        
        .btn:active {
            background: #065f46; /* クリック時はもっと深い緑 */
            transform: translateY(0);
        }
        .info { background: #e9ecef; padding: 15px; border-radius: 4px; margin: 20px 0; }
    </style>
</head>
<body>
    <div class="container">
        <h2>OAuth Client Application</h2>
        <div class="info">
            <p><strong>このアプリケーションは OAuth 2.0 を使用してユーザー認証を行います。</strong></p>
            <p>「ログイン」ボタンをクリックすると、認証サーバーにリダイレクトされます。</p>
        </div>
        
        <h3>OAuth フロー開始</h3>
        <p>認証サーバーでログインするには以下の情報を使用してください：</p>
        <ul>
            <li>ユーザー名: <strong>1</strong></li>
            <li>パスワード: <strong>2</strong></li>
        </ul>
        
        <a href="http://{{.AUTHORIZE_SERVER_DOMAIN}}/authorize?client_id=test-client&redirect_uri=http://localhost:8002/callback&response_type=code&state=test123" class="btn">
            OAuth でログイン
        </a>
        
        <div class="info" style="margin-top: 30px;">
            <h4>フロー説明：</h4>
            <ol>
                <li>認証サーバー ({{.AUTHORIZE_SERVER_DOMAIN}}) にリダイレクト</li>
                <li>ユーザー認証</li>
                <li>認可コードを取得してこのアプリに戻る</li>
                <li>認可コードをアクセストークンに交換</li>
                <li>ユーザー情報を取得</li>
            </ol>
        </div>
    </div>
</body>
</html>`

	tmpl := template.Must(template.New("page").Parse(htmlTemplate))

	w.Header().Set("Content-Type", "text/html")
	data := struct {
		AUTHORIZE_SERVER_DOMAIN string
	}{authorize_server_domain}

	tmpl.Execute(w, data)
}

// コールバックページ（認可コード受け取り用）
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	html := CallbackView()

	t, _ := template.New("callback").Parse(html)
	data := map[string]string{
		"Code":  code,
		"State": state,
	}
	t.Execute(w, data)
}

// トークン交換テスト用ページ
func testTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	code := r.FormValue("code")
	log.Printf("Received code: %s", code)

	// トークンエンドポイントに内部リクエスト
	formData := url.Values{}
	formData.Set("grant_type", "authorization_code")
	formData.Set("code", code)
	formData.Set("client_id", "test-client")
	formData.Set("client_secret", "test-secret")
	formData.Set("redirect_uri", "http://localhost:8002/callback")

	log.Printf("Sending request to token endpoint with data: %s", formData.Encode())

	resp, err := http.PostForm("http://"+authorize_server_docker_domain+"/token", formData)

	if err != nil {
		log.Printf("Token request error: %v", err)
		http.Error(w, "Token request failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	log.Printf("Token response status: %d", resp.StatusCode)

	// レスポンスステータスをチェック
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("Token request failed with status %d: %s", resp.StatusCode, string(body))

		// エラーレスポンスも表示用に保存
		var errorResponse map[string]interface{}
		json.Unmarshal(body, &errorResponse)

		html := TokenExchangeFailedView()

		t, _ := template.New("error").Parse(html)
		data := map[string]string{
			"Status": resp.Status,
			"Error":  string(body),
		}
		t.Execute(w, data)
		return
	}

	var tokenResponse map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&tokenResponse)

	log.Printf("Token response: %+v", tokenResponse)

	// アクセストークンでユーザー情報取得テスト
	var userinfoResponse map[string]interface{}
	if accessToken, ok := tokenResponse["access_token"].(string); ok {
		log.Printf("Access token received: %s", accessToken[:10]+"...")

		req, _ := http.NewRequest("GET", "http://"+authorize_server_docker_domain+"/userinfo", nil)
		req.Header.Set("Authorization", "Bearer "+accessToken)

		client := &http.Client{}
		userinfoResp, err := client.Do(req)
		if err == nil {
			defer userinfoResp.Body.Close()
			log.Printf("Userinfo response status: %d", userinfoResp.StatusCode)
			json.NewDecoder(userinfoResp.Body).Decode(&userinfoResponse)
		} else {
			log.Printf("Userinfo request error: %v", err)
		}
	}

	// 結果表示
	html := OAuthResultView()

	t, _ := template.New("result").Parse(html)

	tokenJSON, _ := json.MarshalIndent(tokenResponse, "", "  ")
	userinfoJSON, _ := json.MarshalIndent(userinfoResponse, "", "  ")

	templateData := map[string]string{
		"TokenResponse":    string(tokenJSON),
		"UserinfoResponse": string(userinfoJSON),
	}
	t.Execute(w, templateData)
}

func main() {
	p := ":" + os.Getenv("GO_PORT")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Health check endpoint hit from %s", r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok", "service": "oauth-client"}`))
	})

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/callback", callbackHandler)
	http.HandleFunc("/test-token", testTokenHandler)

	log.Printf("OAuth Server starting on %v \n", p)
	log.Fatal(http.ListenAndServe(p, nil))
}
