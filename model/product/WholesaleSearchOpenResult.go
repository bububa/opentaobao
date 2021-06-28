package product

// WholesaleSearchOpenResult 
/* model for simplify = false
type WholesaleSearchOpenResult struct {

    // result_memo 返回结果描述，例如success表示成功
    
    ResultMemo   string `json:"result_memo,omitempty"`
    

    // result_code 返回结果码，例如200
    
    ResultCode   int64 `json:"result_code,omitempty"`
    

    // 在线批发产品搜索结果
    
    Result  *struct {
        GoodsSearchResult  *GoodsSearchResult `json:"goods_search_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

// WholesaleSearchOpenResult 
type WholesaleSearchOpenResult struct {

    // result_memo 返回结果描述，例如success表示成功
    ResultMemo   string `json:"result_memo,omitempty"`

    // result_code 返回结果码，例如200
    ResultCode   int64 `json:"result_code,omitempty"`

    // 在线批发产品搜索结果
    Result   *GoodsSearchResult `json:"result,omitempty"`

}
