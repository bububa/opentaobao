package kclub

// AlibabaKclubKcQaGetResult 
/* model for simplify = false
type AlibabaKclubKcQaGetResult struct {

    // 错误信息
    
    Message   string `json:"message,omitempty"`
    

    // 返回数据
    
    Data  *struct {
        KcQaRead  *KcQaRead `json:"kc_qa_read,omitempty"`
    } `json:"data,omitempty"`
    

    // 错误编码
    
    Code   string `json:"code,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlibabaKclubKcQaGetResult 
type AlibabaKclubKcQaGetResult struct {

    // 错误信息
    Message   string `json:"message,omitempty"`

    // 返回数据
    Data   *KcQaRead `json:"data,omitempty"`

    // 错误编码
    Code   string `json:"code,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
