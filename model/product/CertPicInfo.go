package product

// CertPicInfo 
type CertPicInfo struct {

    // 认证类型的数值id
    
    CertType   int64 `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
    

    // 认证图片的url地址
    
    PicUrl   string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
    

}
