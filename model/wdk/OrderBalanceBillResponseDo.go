package wdk

// OrderBalanceBillResponseDo 
type OrderBalanceBillResponseDo struct {

    // orderBalanceBillDOList
    OrderBalanceBillList   []OrderBalanceBillDo `json:"order_balance_bill_list,omitempty"`

    // 是否有下一页0：没有 1：有
    HasNext   string `json:"has_next,omitempty"`

}
