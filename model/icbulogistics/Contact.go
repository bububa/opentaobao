package icbulogistics

// Contact 
type Contact struct {

    // 电话区码
    
    PhoneArea   string `json:"phone_area,omitempty" xml:"phone_area,omitempty"`
    

    // 电话区号
    
    PhoneCode   string `json:"phone_code,omitempty" xml:"phone_code,omitempty"`
    

    // 手机号码
    
    MobileNo   string `json:"mobile_no,omitempty" xml:"mobile_no,omitempty"`
    

    // 邮箱
    
    Email   string `json:"email,omitempty" xml:"email,omitempty"`
    

}
