package tvpay

// ApplyAuthResultDo 
/* model for simplify = false
type ApplyAuthResultDo struct {

    // 授权方式
    
    AuthMode   string `json:"auth_mode,omitempty"`
    

    // 手机号
    
    Mobile   string `json:"mobile,omitempty"`
    

    // 二维码地址
    
    QrCodeUrl   string `json:"qr_code_url,omitempty"`
    

    // qrcode
    
    QrCode   string `json:"qr_code,omitempty"`
    

}
*/

// ApplyAuthResultDo 
type ApplyAuthResultDo struct {

    // 授权方式
    AuthMode   string `json:"auth_mode,omitempty"`

    // 手机号
    Mobile   string `json:"mobile,omitempty"`

    // 二维码地址
    QrCodeUrl   string `json:"qr_code_url,omitempty"`

    // qrcode
    QrCode   string `json:"qr_code,omitempty"`

}
