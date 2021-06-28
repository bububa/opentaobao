package security

// RpidCard 
type RpidCard struct {

    // RPIDCardImage
    
    RpIdcardImage   *RpidCardImage `json:"rp_idcard_image,omitempty" xml:"rp_idcard_image,omitempty"`
    

    // address
    
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    

    // birthDay
    
    BirthDay   string `json:"birth_day,omitempty" xml:"birth_day,omitempty"`
    

    // cardType
    
    CardType   string `json:"card_type,omitempty" xml:"card_type,omitempty"`
    

    // code
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // expiry
    
    Expiry   string `json:"expiry,omitempty" xml:"expiry,omitempty"`
    

    // name
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // sex
    
    Sex   *RpSex `json:"sex,omitempty" xml:"sex,omitempty"`
    

    // urlBackImage
    
    UrlBackImage   string `json:"url_back_image,omitempty" xml:"url_back_image,omitempty"`
    

    // urlFrontImage
    
    UrlFrontImage   string `json:"url_front_image,omitempty" xml:"url_front_image,omitempty"`
    

    // bizErrorCode
    
    BizErrorCode   *RpErrorCode `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
    

}
