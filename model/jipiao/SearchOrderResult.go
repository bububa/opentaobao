package jipiao

// SearchOrderResult 
type SearchOrderResult struct {
    // 淘宝机票订单列表
    OrderIds   []int64 `json:"order_ids,omitempty" xml:"order_ids>int64,omitempty"`
    // 当前淘宝系统设定的搜索结果页大小，即支持一次最多返回订单条数
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 是否还有下一页，即满足搜索条件的订单数是否还有下一个分页
    HasNext   bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}
