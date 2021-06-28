package fenxiao

// Discount 
type Discount struct {

    // 折扣ID
    
    DiscountId   int64 `json:"discount_id,omitempty" xml:"discount_id,omitempty"`
    

    // 折扣名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 折扣详情
    
    Details   []DiscountDetail `json:"details,omitempty" xml:"details,omitempty"`
    

    // 创建时间
    
    Created   string `json:"created,omitempty" xml:"created,omitempty"`
    

    // 修改时间
    
    Modified   string `json:"modified,omitempty" xml:"modified,omitempty"`
    

}
