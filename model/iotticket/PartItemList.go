package iotticket

// PartItemList 
type PartItemList struct {

    // 配件编码
    
    ItemCode   string `json:"item_code,omitempty" xml:"item_code,omitempty"`
    

    // 支付角色：merchant-商家记账；customer-客户付费
    
    PayRole   string `json:"pay_role,omitempty" xml:"pay_role,omitempty"`
    

}
