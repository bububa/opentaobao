package alicom

// OssTokenResponse 
type OssTokenResponse struct {

    // 失效时间
    
    Expiration   string `json:"expiration,omitempty" xml:"expiration,omitempty"`
    

    // token
    
    SecurityToken   string `json:"security_token,omitempty" xml:"security_token,omitempty"`
    

    // accessKeySecret
    
    AccessKeySecret   string `json:"access_key_secret,omitempty" xml:"access_key_secret,omitempty"`
    

    // accessKeyId
    
    AccessKeyId   string `json:"access_key_id,omitempty" xml:"access_key_id,omitempty"`
    

    // status
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

}
