package main

func CallbackView() string {
	return `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>OAuth Callback - 認証完了</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- Material Design Components Web CSS -->
    <link href="https://unpkg.com/material-components-web@latest/dist/material-components-web.min.css" rel="stylesheet">
    <!-- Material Icons -->
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <!-- Roboto Font -->
    <link href="https://fonts.googleapis.com/css?family=Roboto:300,400,500" rel="stylesheet">
    
    <style>
        body {
            font-family: 'Roboto', sans-serif;
            margin: 0;
            padding: 0;
            background: linear-gradient(135deg, #10b981 0%, #059669 100%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        
        .callback-container {
            background: white;
            border-radius: 16px;
            box-shadow: 0 8px 32px rgba(0,0,0,0.1);
            padding: 48px;
            width: 100%;
            max-width: 500px;
            box-sizing: border-box;
            text-align: center;
        }
        
        .success-icon {
            margin-bottom: 24px;
        }
        
        .success-icon .material-icons {
            font-size: 80px;
            color: #4caf50;
            animation: bounce 1s ease-in-out;
        }
        
        @keyframes bounce {
            0%, 20%, 60%, 100% { transform: translateY(0); }
            40% { transform: translateY(-10px); }
            80% { transform: translateY(-5px); }
        }
        
        .callback-title {
            color: #333;
            margin-bottom: 16px;
            font-weight: 400;
            font-size: 28px;
        }
        
        .callback-subtitle {
            color: #666;
            margin-bottom: 32px;
            font-size: 16px;
        }
        
        .info-card {
            background: #f8f9fa;
            border-radius: 12px;
            padding: 24px;
            margin-bottom: 32px;
            border-left: 4px solid #4caf50;
        }
        
        .info-row {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 12px;
            padding: 8px 0;
        }
        
        .info-row:last-child {
            margin-bottom: 0;
        }
        
        .info-label {
            font-weight: 500;
            color: #555;
            display: flex;
            align-items: center;
        }
        
        .info-label .material-icons {
            margin-right: 8px;
            font-size: 20px;
            color: #667eea;
        }
        
        .info-value {
            font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
            background: #e9ecef;
            padding: 6px 12px;
            border-radius: 6px;
            font-size: 14px;
            color: #495057;
            word-break: break-all;
            max-width: 250px;
        }
        
        .action-section {
            margin-top: 32px;
        }
        
        .action-title {
            color: #333;
            margin-bottom: 16px;
            font-size: 20px;
            font-weight: 500;
        }
        
        .action-description {
            color: #666;
            margin-bottom: 24px;
            font-size: 14px;
            line-height: 1.5;
        }

		.mdc-button--raised {
            background-color: #059669 !important; /* 深い緑背景 */
            color: white !important;
        }
        
        .mdc-button--raised:hover {
            background-color: #047857 !important; /* ホバー時はさらに深い緑 */
            box-shadow: 0 4px 12px rgba(5, 150, 105, 0.3) !important;
        }
        
        .mdc-button--raised:active {
            background-color: #065f46 !important; /* クリック時はもっと深い緑 */
        }
        
        .mdc-button--raised:disabled {
            background-color: #6b7280 !important; /* 無効時はグレー */
            color: #9ca3af !important;
        }
        
        .mdc-button {
            width: 100%;
            margin-bottom: 16px;
        }
        
        .next-steps {
            margin-top: 32px;
            padding: 20px;
            background: #e3f2fd;
            border-radius: 8px;
            border-left: 4px solid #2196f3;
            text-align: left;
        }
        
        .next-steps h4 {
            margin: 0 0 12px 0;
            color: #1976d2;
            font-size: 16px;
        }
        
        .next-steps ol {
            margin: 0;
            padding-left: 20px;
            color: #555;
        }
        
        .next-steps li {
            margin-bottom: 8px;
            font-size: 14px;
        }
        
        .loading-spinner {
            display: none;
            margin: 0 auto 16px auto;
            width: 40px;
            height: 40px;
        }
        
        .mdc-circular-progress {
            color: #667eea;
        }
    </style>
</head>
<body>
    <div class="callback-container">
        <div class="success-icon">
            <i class="material-icons">check_circle</i>
        </div>
        
        <h1 class="callback-title">認証コード受信完了</h1>
        <p class="callback-subtitle">OAuth認証の第一段階が正常に完了しました</p>
        
        <div class="info-card">
            <div class="info-row">
                <div class="info-label">
                    <i class="material-icons">vpn_key</i>
                    認証コード
                </div>
                <div class="info-value">{{.Code}}</div>
            </div>
            <div class="info-row">
                <div class="info-label">
                    <i class="material-icons">security</i>
                    状態パラメータ
                </div>
                <div class="info-value">{{.State}}</div>
            </div>
        </div>
        
        <div class="action-section">
            <h3 class="action-title">アクセストークンの取得</h3>
            <p class="action-description">
                認証コードをアクセストークンに交換し、保護されたリソースへのアクセス権限を取得します。
            </p>
            
            <div class="loading-spinner">
                <div class="mdc-circular-progress" role="progressbar" aria-label="Loading..." aria-valuemin="0" aria-valuemax="1">
                    <div class="mdc-circular-progress__determinate-container">
                        <svg class="mdc-circular-progress__determinate-circle-graphic" viewBox="0 0 48 48">
                            <circle class="mdc-circular-progress__determinate-circle" cx="24" cy="24" r="18" stroke-dasharray="113.097" stroke-dashoffset="113.097"></circle>
                        </svg>
                    </div>
                    <div class="mdc-circular-progress__indeterminate-container">
                        <div class="mdc-circular-progress__spinner-layer">
                            <div class="mdc-circular-progress__circle-clipper mdc-circular-progress__circle-left">
                                <svg class="mdc-circular-progress__indeterminate-circle-graphic" viewBox="0 0 48 48">
                                    <circle cx="24" cy="24" r="18" stroke-dasharray="113.097" stroke-dashoffset="56.549"></circle>
                                </svg>
                            </div>
                            <div class="mdc-circular-progress__gap-patch">
                                <svg class="mdc-circular-progress__indeterminate-circle-graphic" viewBox="0 0 48 48">
                                    <circle cx="24" cy="24" r="18" stroke-dasharray="113.097" stroke-dashoffset="56.549"></circle>
                                </svg>
                            </div>
                            <div class="mdc-circular-progress__circle-clipper mdc-circular-progress__circle-right">
                                <svg class="mdc-circular-progress__indeterminate-circle-graphic" viewBox="0 0 48 48">
                                    <circle cx="24" cy="24" r="18" stroke-dasharray="113.097" stroke-dashoffset="56.549"></circle>
                                </svg>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            
            <form action="/test-token" method="POST" onsubmit="showLoading()">
                <input type="hidden" name="code" value="{{.Code}}">
                <button type="submit" class="mdc-button mdc-button--raised">
                    <i class="material-icons mdc-button__icon">swap_horiz</i>
                    <span class="mdc-button__label">アクセストークンに交換</span>
                </button>
            </form>
        </div>
        
        <div class="next-steps">
            <h4><i class="material-icons" style="font-size: 20px; vertical-align: middle; margin-right: 8px;">info</i>次のステップ</h4>
            <ol>
                <li>認証コードをアクセストークンに交換</li>
                <li>トークンを使用してユーザー情報を取得</li>
                <li>セッションの確立と認証完了</li>
            </ol>
        </div>
    </div>
    
    <!-- Material Design Components Web JavaScript -->
    <script src="https://unpkg.com/material-components-web@latest/dist/material-components-web.min.js"></script>
    <script>
        // Material Design Button の初期化
        const buttons = document.querySelectorAll('.mdc-button');
        buttons.forEach(button => {
            mdc.ripple.MDCRipple.attachTo(button);
        });
        
        // ローディング表示
        function showLoading() {
            document.querySelector('.loading-spinner').style.display = 'block';
            document.querySelector('button[type="submit"]').disabled = true;
            document.querySelector('.mdc-button__label').textContent = '処理中...';
        }
        
        // ページロード時のアニメーション
        window.addEventListener('load', function() {
            document.querySelector('.callback-container').style.opacity = '0';
            document.querySelector('.callback-container').style.transform = 'translateY(20px)';
            document.querySelector('.callback-container').style.transition = 'all 0.5s ease-out';
            
            setTimeout(() => {
                document.querySelector('.callback-container').style.opacity = '1';
                document.querySelector('.callback-container').style.transform = 'translateY(0)';
            }, 100);
        });
    </script>
</body>
</html>`
}

func OAuthResultView() string {
	return `
	<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>OAuth Token Result - 認証完了</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- Material Design Components Web CSS -->
    <link href="https://unpkg.com/material-components-web@latest/dist/material-components-web.min.css" rel="stylesheet">
    <!-- Material Icons -->
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <!-- Roboto Font -->
    <link href="https://fonts.googleapis.com/css?family=Roboto:300,400,500" rel="stylesheet">
    <!-- Prism.js for syntax highlighting -->
    <link href="https://cdnjs.cloudflare.com/ajax/libs/prism/1.29.0/themes/prism-tomorrow.min.css" rel="stylesheet">
    
    <style>
        body {
            font-family: 'Roboto', sans-serif;
            margin: 0;
            padding: 20px;
            background: linear-gradient(135deg, #10b981 0%, #059669 100%);
            min-height: 100vh;
        }
        
        .result-container {
            background: white;
            border-radius: 16px;
            box-shadow: 0 8px 32px rgba(0,0,0,0.1);
            padding: 48px;
            max-width: 1000px;
            margin: 0 auto;
            box-sizing: border-box;
        }
        
        .success-header {
            text-align: center;
            margin-bottom: 40px;
        }
        
        .success-icon {
            margin-bottom: 24px;
        }
        
        .success-icon .material-icons {
            font-size: 80px;
            color: #4caf50;
            animation: celebration 2s ease-in-out;
        }
        
        @keyframes celebration {
            0% { transform: scale(0) rotate(0deg); }
            50% { transform: scale(1.2) rotate(180deg); }
            100% { transform: scale(1) rotate(360deg); }
        }
        
        .result-title {
            color: #333;
            margin-bottom: 16px;
            font-weight: 400;
            font-size: 32px;
        }
        
        .result-subtitle {
            color: #666;
            margin-bottom: 0;
            font-size: 18px;
        }
        
        .status-badge {
            display: inline-flex;
            align-items: center;
            background: #e8f5e8;
            color: #2e7d32;
            padding: 8px 16px;
            border-radius: 20px;
            font-size: 14px;
            font-weight: 500;
            margin: 20px 0;
        }
        
        .status-badge .material-icons {
            margin-right: 8px;
            font-size: 18px;
        }
        
        .response-section {
            margin-bottom: 32px;
        }
        
        .section-header {
            display: flex;
            align-items: center;
            margin-bottom: 16px;
            padding-bottom: 12px;
            border-bottom: 2px solid #f0f0f0;
        }
        
        .section-icon {
            background: #667eea;
            color: white;
            width: 40px;
            height: 40px;
            border-radius: 50%;
            display: flex;
            align-items: center;
            justify-content: center;
            margin-right: 16px;
        }
        
        .section-title {
            color: #333;
            font-size: 22px;
            font-weight: 500;
            margin: 0;
        }
        
        .code-container {
            background: #2d3748;
            border-radius: 12px;
            overflow: hidden;
            box-shadow: 0 4px 12px rgba(0,0,0,0.1);
        }
        
        .code-header {
            background: #4a5568;
            padding: 12px 20px;
            display: flex;
            justify-content: between;
            align-items: center;
            border-bottom: 1px solid #2d3748;
        }
        
        .code-language {
            color: #e2e8f0;
            font-size: 12px;
            font-weight: 500;
            text-transform: uppercase;
            letter-spacing: 1px;
        }
        
        .copy-button {
            background: transparent;
            border: 1px solid #667eea;
            color: #667eea;
            padding: 4px 12px;
            border-radius: 4px;
            cursor: pointer;
            font-size: 12px;
            transition: all 0.2s;
            margin-left: auto;
        }
        
        .copy-button:hover {
            background: #667eea;
            color: white;
        }
        
        .code-content {
            padding: 0;
            margin: 0;
            overflow-x: auto;
        }
        
        .code-content pre {
            margin: 0;
            padding: 20px;
            background: transparent;
            color: #e2e8f0;
            font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
            font-size: 14px;
            line-height: 1.5;
            overflow-x: auto;
        }
        
        .action-section {
            text-align: center;
            margin-top: 48px;
            padding-top: 32px;
            border-top: 1px solid #e0e0e0;
        }
        
        .mdc-button {
            margin: 8px;
            min-width: 160px;
        }
        
        .flow-summary {
            background: #f8f9fa;
            border-radius: 12px;
            padding: 24px;
            margin: 32px 0;
            border-left: 4px solid #4caf50;
        }
        
        .flow-summary h4 {
            color: #2e7d32;
            margin: 0 0 16px 0;
            font-size: 18px;
            display: flex;
            align-items: center;
        }
        
        .flow-summary h4 .material-icons {
            margin-right: 12px;
            font-size: 24px;
        }
        
        .flow-steps {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
            gap: 16px;
            margin-top: 16px;
        }
        
        .flow-step {
            background: white;
            padding: 16px;
            border-radius: 8px;
            text-align: center;
            box-shadow: 0 2px 8px rgba(0,0,0,0.05);
        }
        
        .flow-step .material-icons {
            color: #4caf50;
            font-size: 32px;
            margin-bottom: 8px;
        }
        
        .flow-step-title {
            font-weight: 500;
            color: #333;
            margin-bottom: 4px;
            font-size: 14px;
        }
        
        .flow-step-desc {
            color: #666;
            font-size: 12px;
        }
        
        @media (max-width: 768px) {
            .result-container {
                padding: 24px;
                margin: 10px;
            }
            
            .result-title {
                font-size: 24px;
            }
            
            .flow-steps {
                grid-template-columns: 1fr;
            }
        }
		
		.action-section h3 {
            color: #065f46 !important; /* 深い緑でタイトル */
            margin-bottom: 24px;
        }
        
        /* メインボタン（raised）- 深い緑 */
        .mdc-button--raised {
            background-color: #059669 !important;
            color: white !important;
        }
        
        .mdc-button--raised:hover {
            background-color: #047857 !important;
            box-shadow: 0 4px 12px rgba(5, 150, 105, 0.3) !important;
        }
        
        /* アウトラインボタン（outlined）- 緑の枠線 */
        .mdc-button--outlined {
            border-color: #059669 !important;
            color: #059669 !important;
            background-color: transparent !important;
        }
        
        .mdc-button--outlined:hover {
            background-color: #ecfdf5 !important; /* 薄い緑背景 */
            border-color: #047857 !important;
            color: #047857 !important;
        }
        
        .mdc-button--outlined:active {
            background-color: #d1fae5 !important; /* 少し濃い薄緑背景 */
        }
        
        /* ボタンアイコンの色も統一 */
        .mdc-button--raised .material-icons,
        .mdc-button--outlined .material-icons {
            color: inherit !important;
        }
    </style>
</head>
<body>
    <div class="result-container">
        <div class="success-header">
            <div class="success-icon">
                <i class="material-icons">verified</i>
            </div>
            <h1 class="result-title">OAuth 認証完了</h1>
            <p class="result-subtitle">アクセストークンの取得とユーザー情報の取得が正常に完了しました</p>
            <div class="status-badge">
                <i class="material-icons">check_circle</i>
                認証成功
            </div>
        </div>
        
        <div class="flow-summary">
            <h4>
                <i class="material-icons">timeline</i>
                OAuth フロー完了
            </h4>
            <div class="flow-steps">
                <div class="flow-step">
                    <i class="material-icons">login</i>
                    <div class="flow-step-title">認証</div>
                    <div class="flow-step-desc">ユーザー認証完了</div>
                </div>
                <div class="flow-step">
                    <i class="material-icons">vpn_key</i>
                    <div class="flow-step-title">認証コード</div>
                    <div class="flow-step-desc">認証コード取得</div>
                </div>
                <div class="flow-step">
                    <i class="material-icons">swap_horiz</i>
                    <div class="flow-step-title">トークン交換</div>
                    <div class="flow-step-desc">アクセストークン取得</div>
                </div>
                <div class="flow-step">
                    <i class="material-icons">person</i>
                    <div class="flow-step-title">ユーザー情報</div>
                    <div class="flow-step-desc">プロフィール取得</div>
                </div>
            </div>
        </div>
        
        <div class="response-section">
            <div class="section-header">
                <div class="section-icon">
                    <i class="material-icons">token</i>
                </div>
                <h3 class="section-title">アクセストークン レスポンス</h3>
            </div>
            <div class="code-container">
                <div class="code-header">
                    <span class="code-language">JSON Response</span>
                    <button class="copy-button" onclick="copyToClipboard('token-response')">
                        <i class="material-icons" style="font-size: 14px;">content_copy</i> コピー
                    </button>
                </div>
                <div class="code-content">
                    <pre id="token-response" class="language-json">{{.TokenResponse}}</pre>
                </div>
            </div>
        </div>
        
        <div class="response-section">
            <div class="section-header">
                <div class="section-icon">
                    <i class="material-icons">account_circle</i>
                </div>
                <h3 class="section-title">ユーザー情報 レスポンス</h3>
            </div>
            <div class="code-container">
                <div class="code-header">
                    <span class="code-language">JSON Response</span>
                    <button class="copy-button" onclick="copyToClipboard('user-response')">
                        <i class="material-icons" style="font-size: 14px;">content_copy</i> コピー
                    </button>
                </div>
                <div class="code-content">
                    <pre id="user-response" class="language-json">{{.UserinfoResponse}}</pre>
                </div>
            </div>
        </div>
        
        <div class="action-section">
            <h3 style="color: #333; margin-bottom: 24px;">次のアクション</h3>
            <a href="http://localhost:8002/login" class="mdc-button mdc-button--raised">
                <i class="material-icons mdc-button__icon">refresh</i>
                <span class="mdc-button__label">新しい OAuth フローを開始</span>
            </a>
            <a href="http://localhost:8002/health" class="mdc-button mdc-button--outlined">
                <i class="material-icons mdc-button__icon">health_and_safety</i>
                <span class="mdc-button__label">ヘルスチェック</span>
            </a>
        </div>
    </div>
    
    <!-- Material Design Components Web JavaScript -->
    <script src="https://unpkg.com/material-components-web@latest/dist/material-components-web.min.js"></script>
    <!-- Prism.js for syntax highlighting -->
    <script src="https://cdnjs.cloudflare.com/ajax/libs/prism/1.29.0/components/prism-core.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/prism/1.29.0/plugins/autoloader/prism-autoloader.min.js"></script>
    
    <script>
        // Material Design Button の初期化
        const buttons = document.querySelectorAll('.mdc-button');
        buttons.forEach(button => {
            mdc.ripple.MDCRipple.attachTo(button);
        });
        
        // クリップボードにコピー
        function copyToClipboard(elementId) {
            const element = document.getElementById(elementId);
            const text = element.textContent;
            
            navigator.clipboard.writeText(text).then(() => {
                // コピー成功のフィードバック
                const button = event.target.closest('.copy-button');
                const originalText = button.innerHTML;
                button.innerHTML = '<i class="material-icons" style="font-size: 14px;">check</i> コピー済み';
                button.style.background = '#4caf50';
                button.style.color = 'white';
                button.style.borderColor = '#4caf50';
                
                setTimeout(() => {
                    button.innerHTML = originalText;
                    button.style.background = 'transparent';
                    button.style.color = '#667eea';
                    button.style.borderColor = '#667eea';
                }, 2000);
            }).catch(err => {
                console.error('コピーに失敗しました:', err);
            });
        }
        
        // ページロード時のアニメーション
        window.addEventListener('load', function() {
            const container = document.querySelector('.result-container');
            container.style.opacity = '0';
            container.style.transform = 'translateY(30px)';
            container.style.transition = 'all 0.6s ease-out';
            
            setTimeout(() => {
                container.style.opacity = '1';
                container.style.transform = 'translateY(0)';
            }, 200);
            
            // セクションの順次アニメーション
            const sections = document.querySelectorAll('.response-section');
            sections.forEach((section, index) => {
                section.style.opacity = '0';
                section.style.transform = 'translateX(-20px)';
                section.style.transition = 'all 0.5s ease-out';
                
                setTimeout(() => {
                    section.style.opacity = '1';
                    section.style.transform = 'translateX(0)';
                }, 800 + (index * 200));
            });
        });
    </script>
</body>
</html>`
}

func TokenExchangeFailedView() string {
	return `
	<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>OAuth Token Error - エラーが発生しました</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- Material Design Components Web CSS -->
    <link href="https://unpkg.com/material-components-web@latest/dist/material-components-web.min.css" rel="stylesheet">
    <!-- Material Icons -->
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <!-- Roboto Font -->
    <link href="https://fonts.googleapis.com/css?family=Roboto:300,400,500" rel="stylesheet">
    
    <style>
        body {
            font-family: 'Roboto', sans-serif;
            margin: 0;
            padding: 0;
            background: linear-gradient(135deg, #10b981 0%, #059669 100%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        
        .error-container {
            background: white;
            border-radius: 16px;
            box-shadow: 0 8px 32px rgba(0,0,0,0.1);
            padding: 48px;
            width: 100%;
            max-width: 500px;
            box-sizing: border-box;
            text-align: center;
        }
        
        .error-icon {
            margin-bottom: 24px;
        }
        
        .error-icon .material-icons {
            font-size: 80px;
            color: #ef4444;
            animation: shake 0.5s ease-in-out;
        }
        
        @keyframes shake {
            0%, 100% { transform: translateX(0); }
            25% { transform: translateX(-5px); }
            75% { transform: translateX(5px); }
        }
        
        .error-title {
            color: #333;
            margin-bottom: 16px;
            font-weight: 400;
            font-size: 28px;
        }
        
        .error-subtitle {
            color: #666;
            margin-bottom: 32px;
            font-size: 16px;
        }
        
        .error-details {
            background: #fef2f2;
            border-radius: 12px;
            padding: 24px;
            margin-bottom: 32px;
            border-left: 4px solid #ef4444;
            text-align: left;
        }
        
        .error-row {
            display: flex;
            justify-content: space-between;
            align-items: flex-start;
            margin-bottom: 16px;
            flex-wrap: wrap;
        }
        
        .error-row:last-child {
            margin-bottom: 0;
        }
        
        .error-label {
            font-weight: 500;
            color: #dc2626;
            display: flex;
            align-items: center;
            margin-bottom: 8px;
            min-width: 100px;
        }
        
        .error-label .material-icons {
            margin-right: 8px;
            font-size: 20px;
        }
        
        .error-value {
            font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
            background: #fee2e2;
            padding: 8px 12px;
            border-radius: 6px;
            font-size: 14px;
            color: #991b1b;
            word-break: break-word;
            flex: 1;
            min-width: 0;
            margin-left: 16px;
        }
        
        .troubleshooting {
            background: #f0fdf4;
            border-radius: 12px;
            padding: 24px;
            margin-bottom: 32px;
            border-left: 4px solid #10b981;
            text-align: left;
        }
        
        .troubleshooting h4 {
            color: #059669;
            margin: 0 0 16px 0;
            font-size: 18px;
            display: flex;
            align-items: center;
        }
        
        .troubleshooting h4 .material-icons {
            margin-right: 12px;
            font-size: 24px;
        }
        
        .troubleshooting ul {
            margin: 0;
            padding-left: 20px;
            color: #065f46;
        }
        
        .troubleshooting li {
            margin-bottom: 8px;
            font-size: 14px;
            line-height: 1.5;
        }
        
        .action-section {
            margin-top: 32px;
        }
        
        .action-title {
            color: #333;
            margin-bottom: 16px;
            font-size: 20px;
            font-weight: 500;
        }
        
        .action-description {
            color: #065f46;
            margin-bottom: 24px;
            font-size: 14px;
            line-height: 1.5;
        }
        
        .mdc-button {
            width: 100%;
            margin-bottom: 12px;
        }
        
        .mdc-button--raised {
            background-color: #059669 !important;
            color: white !important;
        }
        
        .mdc-button--raised:hover {
            background-color: #047857 !important;
            box-shadow: 0 4px 12px rgba(5, 150, 105, 0.3) !important;
        }
        
        .mdc-button--outlined {
            border-color: #059669 !important;
            color: #059669 !important;
        }
        
        .mdc-button--outlined:hover {
            background-color: #ecfdf5 !important;
            border-color: #047857 !important;
            color: #047857 !important;
        }
        
        .status-info {
            display: inline-flex;
            align-items: center;
            background: #fef2f2;
            color: #dc2626;
            padding: 8px 16px;
            border-radius: 20px;
            font-size: 14px;
            font-weight: 500;
            margin: 20px 0;
        }
        
        .status-info .material-icons {
            margin-right: 8px;
            font-size: 18px;
        }
        
        @media (max-width: 768px) {
            .error-container {
                padding: 24px;
                margin: 10px;
            }
            
            .error-title {
                font-size: 24px;
            }
            
            .error-row {
                flex-direction: column;
            }
            
            .error-value {
                margin-left: 0;
                margin-top: 8px;
            }
        }
    </style>
</head>
<body>
    <div class="error-container">
        <div class="error-icon">
            <i class="material-icons">error</i>
        </div>
        
        <h1 class="error-title">トークン交換に失敗しました</h1>
        <p class="error-subtitle">OAuth認証プロセス中にエラーが発生しました</p>
        
        <div class="status-info">
            <i class="material-icons">warning</i>
            認証エラー
        </div>
        
        <div class="error-details">
            <div class="error-row">
                <div class="error-label">
                    <i class="material-icons">info</i>
                    ステータス
                </div>
                <div class="error-value">{{.Status}}</div>
            </div>
            <div class="error-row">
                <div class="error-label">
                    <i class="material-icons">bug_report</i>
                    エラー詳細
                </div>
                <div class="error-value">{{.Error}}</div>
            </div>
        </div>
        
        <div class="troubleshooting">
            <h4>
                <i class="material-icons">build</i>
                トラブルシューティング
            </h4>
            <ul>
                <li>デモのため認証コード毎リソースアクセス回数は1回です。</li>
                <li>認証コードの有効期限が切れている可能性があります（10分間有効）</li>
                <li>クライアントIDまたはシークレットが正しくない可能性があります</li>
                <li>リダイレクトURIが一致していない可能性があります</li>
                <li>OAuthサーバーとの通信に問題がある可能性があります</li>
                <li>認証コードが既に使用済みの可能性があります</li>
            </ul>
        </div>
        
        <div class="action-section">
            <h3 class="action-title">解決方法</h3>
            <p class="action-description">
                新しいOAuthフローを開始して、再度認証を試行してください。
            </p>
            
            <a href="http://localhost:8001/authorize?client_id=test-client&redirect_uri=http://localhost:8002/callback&response_type=code&state=test123" class="mdc-button mdc-button--raised">
                <i class="material-icons mdc-button__icon">refresh</i>
                <span class="mdc-button__label">新しい OAuth フローを開始</span>
            </a>
            
            <a href="http://localhost:8002/" class="mdc-button mdc-button--outlined">
                <i class="material-icons mdc-button__icon">home</i>
                <span class="mdc-button__label">ヘルスチェック</span>
            </a>
        </div>
    </div>
    
    <!-- Material Design Components Web JavaScript -->
    <script src="https://unpkg.com/material-components-web@latest/dist/material-components-web.min.js"></script>
    <script>
        // Material Design Button の初期化
        const buttons = document.querySelectorAll('.mdc-button');
        buttons.forEach(button => {
            mdc.ripple.MDCRipple.attachTo(button);
        });
        
        // ページロード時のアニメーション
        window.addEventListener('load', function() {
            const container = document.querySelector('.error-container');
            container.style.opacity = '0';
            container.style.transform = 'translateY(20px)';
            container.style.transition = 'all 0.5s ease-out';
            
            setTimeout(() => {
                container.style.opacity = '1';
                container.style.transform = 'translateY(0)';
            }, 100);
        });
    </script>
</body>
</html>
	`
}
