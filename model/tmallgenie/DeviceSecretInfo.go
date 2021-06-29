package tmallgenie

// DeviceSecretInfo 
type DeviceSecretInfo struct {

    // 错误原因
    
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
    

    // 获取秘钥是否成功
    
    Sucess   bool `json:"sucess,omitempty" xml:"sucess,omitempty"`
    

    // 秘钥
    
    Secret   string `json:"secret,omitempty" xml:"secret,omitempty"`
    

    // mac地址
    
    Mac   string `json:"mac,omitempty" xml:"mac,omitempty"`
    

    // md5
    
    Md5   string `json:"md5,omitempty" xml:"md5,omitempty"`
    

}
