package omniorder

// TaobaoOmniitemClassifyQueryResult 
type TaobaoOmniitemClassifyQueryResult struct {

    // 提示信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 业务数据
    
    Datas   []Classify `json:"datas,omitempty" xml:"datas,omitempty"`
    

    // 状态码
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

}
