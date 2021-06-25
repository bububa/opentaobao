package alicom

// IdCardModel 
type IdCardModel struct {

    // 身份证反面照
    BackCertPic   string `json:"back_cert_pic,omitempty"`

    // 证件失效时间
    CardExpireDate   string `json:"card_expire_date,omitempty"`

    // 证件号码
    CardNumber   string `json:"card_number,omitempty"`

    // 证件类型
    CardType   string `json:"card_type,omitempty"`

    // 身份证正面照
    FaceCertPic   string `json:"face_cert_pic,omitempty"`

    // 手持身份证照
    HoldCertPic   string `json:"hold_cert_pic,omitempty"`

    // 证件名称
    Name   string `json:"name,omitempty"`

}
