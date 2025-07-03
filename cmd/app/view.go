package main

func AuthorizeView() string {
	return `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>OAuth Login</title>
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
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        
        .login-container {
            background: white;
            border-radius: 16px;
            box-shadow: 0 8px 32px rgba(0,0,0,0.1);
            padding: 48px;
            width: 100%;
            max-width: 400px;
            box-sizing: border-box;
        }
        
        .app-icon {
            text-align: center;
            margin-bottom: 32px;
        }
        
        .app-icon .material-icons {
            font-size: 64px;
            color: #667eea;
        }
        
        .login-title {
            text-align: center;
            color: #333;
            margin-bottom: 8px;
            font-weight: 400;
            font-size: 28px;
        }
        
        .login-subtitle {
            text-align: center;
            color: #666;
            margin-bottom: 32px;
            font-size: 14px;
        }
        
        .client-info {
            background: #f5f5f5;
            border-radius: 8px;
            padding: 16px;
            margin-bottom: 32px;
            text-align: center;
        }
        
        .client-info .material-icons {
            color: #667eea;
            margin-right: 8px;
            vertical-align: middle;
        }
        
        .mdc-text-field {
            width: 100%;
            margin-bottom: 24px;
        }
        
        .mdc-button {
            width: 100%;
            margin-top: 16px;
        }
        
        .security-notice {
            margin-top: 24px;
            padding: 16px;
            background: #fff3cd;
            border-radius: 8px;
            border-left: 4px solid #ffc107;
            font-size: 12px;
            color: #856404;
        }
    </style>
</head>
<body>
    <div class="login-container">
        <div class="app-icon">
            <i class="material-icons">security</i>
        </div>
        
        <h1 class="login-title">認証が必要です</h1>
        <p class="login-subtitle">アカウントにアクセスするための認証を行います</p>
        
        <div class="client-info">
            <i class="material-icons">apps</i>
            <strong>{{.ClientID}}</strong> があなたのアカウントへのアクセスを要求しています
        </div>
        
        <form method="POST">
            <input type="hidden" name="client_id" value="{{.ClientID}}">
            <input type="hidden" name="redirect_uri" value="{{.RedirectURI}}">
            <input type="hidden" name="response_type" value="{{.ResponseType}}">
            <input type="hidden" name="state" value="{{.State}}">
            
            <div class="mdc-text-field mdc-text-field--outlined">
                <input type="text" name="username" required class="mdc-text-field__input" id="username">
                <div class="mdc-notched-outline">
                    <div class="mdc-notched-outline__leading"></div>
                    <div class="mdc-notched-outline__notch">
                        <label for="username" class="mdc-floating-label">ユーザー名</label>
                    </div>
                    <div class="mdc-notched-outline__trailing"></div>
                </div>
            </div>
            
            <div class="mdc-text-field mdc-text-field--outlined">
                <input type="password" name="password" required class="mdc-text-field__input" id="password">
                <div class="mdc-notched-outline">
                    <div class="mdc-notched-outline__leading"></div>
                    <div class="mdc-notched-outline__notch">
                        <label for="password" class="mdc-floating-label">パスワード</label>
                    </div>
                    <div class="mdc-notched-outline__trailing"></div>
                </div>
            </div>
            
            <button type="submit" class="mdc-button mdc-button--raised">
                <span class="mdc-button__label">認証してアクセスを許可</span>
            </button>
        </form>
        
        <div class="security-notice">
            <i class="material-icons" style="font-size: 16px; vertical-align: middle; margin-right: 8px;">info</i>
            この認証は安全な接続を通じて行われます。アクセス許可後、アプリケーションは限定された情報のみにアクセスできます。
        </div>
    </div>
    
    <!-- Material Design Components Web JavaScript -->
    <script src="https://unpkg.com/material-components-web@latest/dist/material-components-web.min.js"></script>
    <script>
        // Material Design Text Field の初期化
        const textFields = document.querySelectorAll('.mdc-text-field');
        textFields.forEach(textField => {
            mdc.textField.MDCTextField.attachTo(textField);
        });
        
        // Material Design Button の初期化
        const buttons = document.querySelectorAll('.mdc-button');
        buttons.forEach(button => {
            mdc.ripple.MDCRipple.attachTo(button);
        });
    </script>
</body>
</html>`
}
