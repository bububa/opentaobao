package product

// TaobaoBanamadpcItemEditRenderApiResult 
/* model for simplify = false
type TaobaoBanamadpcItemEditRenderApiResult struct {

    // 错误信息
    
    ErMsg   string `json:"er_msg,omitempty"`
    

    // 错误码
    
    ErCode   string `json:"er_code,omitempty"`
    

    // 商品标题，价格，图片等商品数据的schema xml
    
    Result   string `json:"result,omitempty"`
    

    // 成功
    
    Error   bool `json:"error,omitempty"`
    

}
*/

// TaobaoBanamadpcItemEditRenderApiResult 
type TaobaoBanamadpcItemEditRenderApiResult struct {

    // 错误信息
    ErMsg   string `json:"er_msg,omitempty"`

    // 错误码
    ErCode   string `json:"er_code,omitempty"`

    // 商品标题，价格，图片等商品数据的schema xml
    Result   string `json:"result,omitempty"`

    // 成功
    Error   bool `json:"error,omitempty"`

}
