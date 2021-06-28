package wdk

// OrderOperateResult 
/* model for simplify = false
type OrderOperateResult struct {

    // 盒马主单号
    
    BizOrderId   string `json:"biz_order_id,omitempty"`
    

    // 外部主单号
    
    OutOrderId   string `json:"out_order_id,omitempty"`
    

    // 子单列表信息
    
    SubOrderResults  struct {
        SubOrder  []SubOrder `json:"sub_order,omitempty"`
    } `json:"sub_order_results,omitempty"`
    

}
*/

// OrderOperateResult 
type OrderOperateResult struct {

    // 盒马主单号
    BizOrderId   string `json:"biz_order_id,omitempty"`

    // 外部主单号
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 子单列表信息
    SubOrderResults   []SubOrder `json:"sub_order_results,omitempty"`

}
