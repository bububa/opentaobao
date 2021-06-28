package product

// TaobaoBanamadpcItemRenderApiResult 
/* model for simplify = false
type TaobaoBanamadpcItemRenderApiResult struct {

    // 错误信息
    
    ErMsg   string `json:"er_msg,omitempty"`
    

    // 错误码
    
    ErCode   string `json:"er_code,omitempty"`
    

    // 返回类目新发下需要填写的商品信息
    
    Result   string `json:"result,omitempty"`
    

    // 错误
    
    Error   bool `json:"error,omitempty"`
    

}
*/

// TaobaoBanamadpcItemRenderApiResult 
type TaobaoBanamadpcItemRenderApiResult struct {

    // 错误信息
    ErMsg   string `json:"er_msg,omitempty"`

    // 错误码
    ErCode   string `json:"er_code,omitempty"`

    // 返回类目新发下需要填写的商品信息
    Result   string `json:"result,omitempty"`

    // 错误
    Error   bool `json:"error,omitempty"`

}
