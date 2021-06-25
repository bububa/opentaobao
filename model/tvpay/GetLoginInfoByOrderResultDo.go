package tvpay

// GetLoginInfoByOrderResultDo 
type GetLoginInfoByOrderResultDo struct {

    // 登陆信息，json
    AccessData   string `json:"access_data,omitempty"`

    // 是否有登陆信息
    HasLoginInfo   bool `json:"has_login_info,omitempty"`

}
