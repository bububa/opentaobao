package alitrippoi

// BaseResult 
type BaseResult struct {

    // 返回的数据实体
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    

    // 是否执行成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 时间
    
    ServerTime   int64 `json:"server_time,omitempty" xml:"server_time,omitempty"`
    

    // 标题
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 错误信息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    

    // 错误码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    

}
