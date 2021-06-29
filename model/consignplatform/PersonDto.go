package consignplatform

// PersonDto 
type PersonDto struct {

    // 收件固定电话
    
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    

    // 收件电话
    
    MobilePhone   string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
    

    // 收件人
    
    UserName   string `json:"user_name,omitempty" xml:"user_name,omitempty"`
    

}
