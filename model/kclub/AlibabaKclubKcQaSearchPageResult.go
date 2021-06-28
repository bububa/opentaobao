package kclub

// AlibabaKclubKcQaSearchPageResult 
/* model for simplify = false
type AlibabaKclubKcQaSearchPageResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 错误编码
    
    Code   string `json:"code,omitempty"`
    

    // 错误信息
    
    Message   string `json:"message,omitempty"`
    

    // 分页数据
    
    Data  *struct {
        Paging  *Paging `json:"paging,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

// AlibabaKclubKcQaSearchPageResult 
type AlibabaKclubKcQaSearchPageResult struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 错误编码
    Code   string `json:"code,omitempty"`

    // 错误信息
    Message   string `json:"message,omitempty"`

    // 分页数据
    Data   *Paging `json:"data,omitempty"`

}
