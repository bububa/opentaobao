package security

// RpidCardBo 
type RpidCardBo struct {

    // 类型
    CardType   string `json:"card_type,omitempty"`

    // 身份证背面照的URL地址
    UrlBackImage   string `json:"url_back_image,omitempty"`

    // 身份证正面照的URL地址
    UrlFrontImage   string `json:"url_front_image,omitempty"`

    // 身份证上生日
    BirthDay   string `json:"birth_day,omitempty"`

    // 身份证住址
    Address   string `json:"address,omitempty"`

    // 性别
    Sex   *RpSex `json:"sex,omitempty"`

    // 身份证上的名字
    Name   string `json:"name,omitempty"`

    // 公民身份证号码
    Code   string `json:"code,omitempty"`

    // RPIDCardImage
    RPIDCardImage   *RpidCardImage `json:"r_p_i_d_card_image,omitempty"`

}