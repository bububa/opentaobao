package wdk

// CreateReverseResponse 
/* model for simplify = false
type CreateReverseResponse struct {

    // tp单号
    
    OutOrderId   string `json:"out_order_id,omitempty"`
    

    // 请求id
    
    RequestId   string `json:"request_id,omitempty"`
    

    // 逆向单ids
    
    ReverseIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"reverse_ids,omitempty"`
    

    // 门店id
    
    StoreId   string `json:"store_id,omitempty"`
    

    // 外部单号
    
    OutBizOrderIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"out_biz_order_ids,omitempty"`
    

}
*/

// CreateReverseResponse 
type CreateReverseResponse struct {

    // tp单号
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 请求id
    RequestId   string `json:"request_id,omitempty"`

    // 逆向单ids
    ReverseIds   []int64 `json:"reverse_ids,omitempty"`

    // 门店id
    StoreId   string `json:"store_id,omitempty"`

    // 外部单号
    OutBizOrderIds   []string `json:"out_biz_order_ids,omitempty"`

}
