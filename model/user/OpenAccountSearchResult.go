package user

// OpenAccountSearchResult 
type OpenAccountSearchResult struct {

    // OpenAccount的列表
    
    Datas   []OpenAccount `json:"datas,omitempty" xml:"datas,omitempty"`
    

    // 查询是否成功，成功返回时有可能数据为空
    
    Successful   bool `json:"successful,omitempty" xml:"successful,omitempty"`
    

    // 状态码
    
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    

    // 状态信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

}
