package product

// WholesaleGoodsOpenResult 
type WholesaleGoodsOpenResult struct {

    // result_memo 返回结果描述，例如success表示成功
    
    ResultMemo   string `json:"result_memo,omitempty" xml:"result_memo,omitempty"`
    

    // result_code 返回码，例如200表示成功
    
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
    

    // 产品详细信息
    
    Result   *GoodsDetail `json:"result,omitempty" xml:"result,omitempty"`
    

}
