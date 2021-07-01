package koubeimall

// TaobaoKoubeiMallCommonItemShelfPageResult 
type TaobaoKoubeiMallCommonItemShelfPageResult struct {
    // API请求全链路追踪ID
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    // 门店货架商品列表信息
    Data   *ShopShelfDto `json:"data,omitempty" xml:"data,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误信息，success=false时，返回相关错误信息
    Error   *TribeError `json:"error,omitempty" xml:"error,omitempty"`
}
