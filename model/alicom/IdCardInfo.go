package alicom

// IdCardInfo 
/* model for simplify = false
type IdCardInfo struct {

    // 证件号码
    
    CardNumber   string `json:"card_number,omitempty"`
    

    // 证件类型
    
    CardType   string `json:"card_type,omitempty"`
    

    // 证件名称
    
    Name   string `json:"name,omitempty"`
    

    // 身份证反面照
    
    BackCertPic   string `json:"back_cert_pic,omitempty"`
    

    // 身份证正面照
    
    FaceCertPic   string `json:"face_cert_pic,omitempty"`
    

    // 手持证件照
    
    HoldCertPic   string `json:"hold_cert_pic,omitempty"`
    

}
*/

// IdCardInfo 
type IdCardInfo struct {

    // 证件号码
    CardNumber   string `json:"card_number,omitempty"`

    // 证件类型
    CardType   string `json:"card_type,omitempty"`

    // 证件名称
    Name   string `json:"name,omitempty"`

    // 身份证反面照
    BackCertPic   string `json:"back_cert_pic,omitempty"`

    // 身份证正面照
    FaceCertPic   string `json:"face_cert_pic,omitempty"`

    // 手持证件照
    HoldCertPic   string `json:"hold_cert_pic,omitempty"`

}
