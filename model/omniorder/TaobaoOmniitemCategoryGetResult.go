package omniorder

// TaobaoOmniitemCategoryGetResult 
type TaobaoOmniitemCategoryGetResult struct {

    // 错误码
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // data
    
    Datas   []OmniItemCategoryDto `json:"datas,omitempty" xml:"datas,omitempty"`
    

    // 错误信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
