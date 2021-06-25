package tvpay

// SdkGlobalConfigDo 
type SdkGlobalConfigDo struct {

    // 是否询问用户要自动登录
    AskBeforeAutoLogin   bool `json:"ask_before_auto_login,omitempty"`

    // 如果询问，是否默认选中自动登录
    AutoLoginChecked   bool `json:"auto_login_checked,omitempty"`

    // 询问对话框样式
    AutoLoginDialogPattern   int64 `json:"auto_login_dialog_pattern,omitempty"`

    // 是否默认展示授权二维码
    ShowAuthCodeByDefault   bool `json:"show_auth_code_by_default,omitempty"`

    // h5路由属性
    H5ProxyUrl   string `json:"h5_proxy_url,omitempty"`

}
