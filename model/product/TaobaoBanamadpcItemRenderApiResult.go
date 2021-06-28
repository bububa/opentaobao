package product

// TaobaoBanamadpcItemRenderApiResult 
type TaobaoBanamadpcItemRenderApiResult struct {

    // 错误信息
    
    ErMsg   string `json:"er_msg,omitempty" xml:"er_msg,omitempty"`
    

    // 错误码
    
    ErCode   string `json:"er_code,omitempty" xml:"er_code,omitempty"`
    

    // 返回类目新发下需要填写的商品信息
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
    

    // 错误
    
    Error   bool `json:"error,omitempty" xml:"error,omitempty"`
    

}
