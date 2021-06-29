package omniorder

// TaobaoOmniitemClassifyItemQueryResult 
type TaobaoOmniitemClassifyItemQueryResult struct {

    // 提示信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 商品ID集合
    
    Datas   []string `json:"datas,omitempty" xml:"datas>string,omitempty"`
    

    // 状态码
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

}
