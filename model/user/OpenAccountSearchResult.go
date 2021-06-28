package user

// OpenAccountSearchResult 
/* model for simplify = false
type OpenAccountSearchResult struct {

    // OpenAccount的列表
    
    Datas  struct {
        OpenAccount  []OpenAccount `json:"open_account,omitempty"`
    } `json:"datas,omitempty"`
    

    // 查询是否成功，成功返回时有可能数据为空
    
    Successful   bool `json:"successful,omitempty"`
    

    // 状态码
    
    Code   int64 `json:"code,omitempty"`
    

    // 状态信息
    
    Message   string `json:"message,omitempty"`
    

}
*/

// OpenAccountSearchResult 
type OpenAccountSearchResult struct {

    // OpenAccount的列表
    Datas   []OpenAccount `json:"datas,omitempty"`

    // 查询是否成功，成功返回时有可能数据为空
    Successful   bool `json:"successful,omitempty"`

    // 状态码
    Code   int64 `json:"code,omitempty"`

    // 状态信息
    Message   string `json:"message,omitempty"`

}
