package wdk

// WdkOpenOrderFinanceBillQueryResult 
type WdkOpenOrderFinanceBillQueryResult struct {

    // 总数量，只在查询第一页时返回
    TotalNumber   int64 `json:"total_number,omitempty"`

    // 下一页查询的入参，当为-1时表示没有更多数据
    NextId   int64 `json:"next_id,omitempty"`

    // 结果信息
    ReturnMsg   string `json:"return_msg,omitempty"`

    // 结果码
    ReturnCode   string `json:"return_code,omitempty"`

    // 成功或失败，调用方需要根据该状态判断是否成功
    Success   bool `json:"success,omitempty"`

    // 账单列表
    Bills   []WdkOpenOrderFinanceBillDo `json:"bills,omitempty"`

}
