package fivee

// BatchProduct 
type BatchProduct struct {

    // 到期日期
    
    DueDate   string `json:"due_date,omitempty" xml:"due_date,omitempty"`
    

    // 原产国
    
    OriginCountry   string `json:"origin_country,omitempty" xml:"origin_country,omitempty"`
    

    // 生产日期
    
    ProduceDate   string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
    

}
