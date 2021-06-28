package security

// RpidCard 
/* model for simplify = false
type RpidCard struct {

    // RPIDCardImage
    
    RpIdcardImage  *struct {
        RpidCardImage  *RpidCardImage `json:"rpid_card_image,omitempty"`
    } `json:"rp_idcard_image,omitempty"`
    

    // address
    
    Address   string `json:"address,omitempty"`
    

    // birthDay
    
    BirthDay   string `json:"birth_day,omitempty"`
    

    // cardType
    
    CardType   string `json:"card_type,omitempty"`
    

    // code
    
    Code   string `json:"code,omitempty"`
    

    // expiry
    
    Expiry   string `json:"expiry,omitempty"`
    

    // name
    
    Name   string `json:"name,omitempty"`
    

    // sex
    
    Sex  *struct {
        RpSex  *RpSex `json:"rp_sex,omitempty"`
    } `json:"sex,omitempty"`
    

    // urlBackImage
    
    UrlBackImage   string `json:"url_back_image,omitempty"`
    

    // urlFrontImage
    
    UrlFrontImage   string `json:"url_front_image,omitempty"`
    

    // bizErrorCode
    
    BizErrorCode  *struct {
        RpErrorCode  *RpErrorCode `json:"rp_error_code,omitempty"`
    } `json:"biz_error_code,omitempty"`
    

}
*/

// RpidCard 
type RpidCard struct {

    // RPIDCardImage
    RpIdcardImage   *RpidCardImage `json:"rp_idcard_image,omitempty"`

    // address
    Address   string `json:"address,omitempty"`

    // birthDay
    BirthDay   string `json:"birth_day,omitempty"`

    // cardType
    CardType   string `json:"card_type,omitempty"`

    // code
    Code   string `json:"code,omitempty"`

    // expiry
    Expiry   string `json:"expiry,omitempty"`

    // name
    Name   string `json:"name,omitempty"`

    // sex
    Sex   *RpSex `json:"sex,omitempty"`

    // urlBackImage
    UrlBackImage   string `json:"url_back_image,omitempty"`

    // urlFrontImage
    UrlFrontImage   string `json:"url_front_image,omitempty"`

    // bizErrorCode
    BizErrorCode   *RpErrorCode `json:"biz_error_code,omitempty"`

}
