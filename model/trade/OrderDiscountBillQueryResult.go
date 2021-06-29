package trade

// OrderDiscountBillQueryResult 
type OrderDiscountBillQueryResult struct {
    // 下一页查询的入参，当为-1时表示没有更多数据
    NextId   int64 `json:"next_id,omitempty" xml:"next_id,omitempty"`
    // 结果码
    ReturnCode   string `json:"return_code,omitempty" xml:"return_code,omitempty"`
    // 结果文案
    ReturnMsg   string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
    // 业务请求成功与否
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 总数量，只在查询第一页时返回
    TotalNumber   int64 `json:"total_number,omitempty" xml:"total_number,omitempty"`
    // 账单列表
    DiscountBills   []OrderDiscountBillBo `json:"discount_bills,omitempty" xml:"discount_bills>order_discount_bill_bo,omitempty"`
}
