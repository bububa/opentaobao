package tvpay

// SdkAppConfigDo 
/* model for simplify = false
type SdkAppConfigDo struct {

    // 是否上传日志
    
    EnableUploadLog   bool `json:"enable_upload_log,omitempty"`
    

    // 是否开启友盟
    
    EnableYoument   bool `json:"enable_youment,omitempty"`
    

    // 自定义属性
    
    ConfigProps   string `json:"config_props,omitempty"`
    

    // 是否自动登入
    
    EnableAutoLogin   bool `json:"enable_auto_login,omitempty"`
    

}
*/

// SdkAppConfigDo 
type SdkAppConfigDo struct {

    // 是否上传日志
    EnableUploadLog   bool `json:"enable_upload_log,omitempty"`

    // 是否开启友盟
    EnableYoument   bool `json:"enable_youment,omitempty"`

    // 自定义属性
    ConfigProps   string `json:"config_props,omitempty"`

    // 是否自动登入
    EnableAutoLogin   bool `json:"enable_auto_login,omitempty"`

}
