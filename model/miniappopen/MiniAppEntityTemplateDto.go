package miniappopen

// MiniAppEntityTemplateDTO 
type MiniAppEntityTemplateDTO struct {
    // 小程序描述
    AppDescription   string `json:"app_description,omitempty" xml:"app_description,omitempty"`
    // 小程序名称
    AppName   string `json:"app_name,omitempty" xml:"app_name,omitempty"`
    // appkey
    Appkey   string `json:"appkey,omitempty" xml:"appkey,omitempty"`
    // 小程序app_id
    AppId   string `json:"app_id,omitempty" xml:"app_id,omitempty"`
    // 小程序icon
    AppIcon   string `json:"app_icon,omitempty" xml:"app_icon,omitempty"`
    // 线上正式版本的链接，所有消费者可访问。
    OnlineUrl   string `json:"online_url,omitempty" xml:"online_url,omitempty"`
    // 应用id
    Id   string `json:"id,omitempty" xml:"id,omitempty"`
    // 线上版本
    OnlineVersion   string `json:"online_version,omitempty" xml:"online_version,omitempty"`
    // 线上码
    OnlineCode   string `json:"online_code,omitempty" xml:"online_code,omitempty"`
}
