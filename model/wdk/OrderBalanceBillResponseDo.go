package wdk

// OrderBalanceBillResponseDO 
type OrderBalanceBillResponseDO struct {
    // orderBalanceBillDOList
    OrderBalanceBillList   []OrderBalanceBillDO `json:"order_balance_bill_list,omitempty" xml:"order_balance_bill_list>order_balance_bill_do,omitempty"`
    // 是否有下一页0：没有 1：有
    HasNext   string `json:"has_next,omitempty" xml:"has_next,omitempty"`
}
