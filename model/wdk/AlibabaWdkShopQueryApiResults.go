package wdk

// AlibabaWdkShopQueryApiResults 
/* model for simplify = false
type AlibabaWdkShopQueryApiResults struct {

    // 返回门店信息列表
    
    Models  struct {
        ShopDo  []ShopDo `json:"shop_do,omitempty"`
    } `json:"models,omitempty"`
    

    // 异常编码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 异常信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // true
    
    Success   string `json:"success,omitempty"`
    

}
*/

// AlibabaWdkShopQueryApiResults 
type AlibabaWdkShopQueryApiResults struct {

    // 返回门店信息列表
    Models   []ShopDo `json:"models,omitempty"`

    // 异常编码
    ErrCode   string `json:"err_code,omitempty"`

    // 异常信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // true
    Success   string `json:"success,omitempty"`

}
