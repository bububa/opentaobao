package tmallhk

// CtsDelivery 
type CtsDelivery struct {

    // 快递公司名称
    
    CompanyName   string `json:"company_name,omitempty" xml:"company_name,omitempty"`
    

    // 快递单号
    
    DeliveryNo   string `json:"delivery_no,omitempty" xml:"delivery_no,omitempty"`
    

    // 发件时间
    
    DeliveryTime   string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
    

}
