package wdk

// OrderBalanceBillResponseDo 
/* model for simplify = false
type OrderBalanceBillResponseDo struct {

    // orderBalanceBillDOList
    
    OrderBalanceBillList  struct {
        OrderBalanceBillDo  []OrderBalanceBillDo `json:"order_balance_bill_do,omitempty"`
    } `json:"order_balance_bill_list,omitempty"`
    

    // 是否有下一页0：没有 1：有
    
    HasNext   string `json:"has_next,omitempty"`
    

}
*/

// OrderBalanceBillResponseDo 
type OrderBalanceBillResponseDo struct {

    // orderBalanceBillDOList
    OrderBalanceBillList   []OrderBalanceBillDo `json:"order_balance_bill_list,omitempty"`

    // 是否有下一页0：没有 1：有
    HasNext   string `json:"has_next,omitempty"`

}
