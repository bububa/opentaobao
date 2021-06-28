package waybill

// Address 
type Address struct {

    // 市
    
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    

    // 详细地址
    
    Detail   string `json:"detail,omitempty" xml:"detail,omitempty"`
    

    // 区
    
    District   string `json:"district,omitempty" xml:"district,omitempty"`
    

    // 省
    
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
    

    // 街道
    
    Town   string `json:"town,omitempty" xml:"town,omitempty"`
    

}
