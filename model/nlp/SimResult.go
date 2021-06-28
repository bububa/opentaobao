package nlp

// SimResult 
/* model for simplify = false
type SimResult struct {

    // 返回文本处理内容
    
    TopResult   string `json:"top_result,omitempty"`
    

    // 返回结果为true则运行成功，为false则运行失败
    
    TopStatus   bool `json:"top_status,omitempty"`
    

}
*/

// SimResult 
type SimResult struct {

    // 返回文本处理内容
    TopResult   string `json:"top_result,omitempty"`

    // 返回结果为true则运行成功，为false则运行失败
    TopStatus   bool `json:"top_status,omitempty"`

}
