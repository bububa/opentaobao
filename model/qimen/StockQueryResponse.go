package qimen

// StockQueryResponse 
/* model for simplify = false
type StockQueryResponse struct {

    // 响应结果:success|failure
    
    Flag   string `json:"flag,omitempty"`
    

    // 响应码
    
    Code   string `json:"code,omitempty"`
    

    // 响应信息
    
    Message   string `json:"message,omitempty"`
    

    // 总数
    
    TotalCount   int64 `json:"totalCount,omitempty"`
    

    // 商品的库存信息列表
    
    Items  struct {
        Item  []Item `json:"item,omitempty"`
    } `json:"items,omitempty"`
    

}
*/

// StockQueryResponse 
type StockQueryResponse struct {

    // 响应结果:success|failure
    Flag   string `json:"flag,omitempty"`

    // 响应码
    Code   string `json:"code,omitempty"`

    // 响应信息
    Message   string `json:"message,omitempty"`

    // 总数
    TotalCount   int64 `json:"totalCount,omitempty"`

    // 商品的库存信息列表
    Items   []Item `json:"items,omitempty"`

}
