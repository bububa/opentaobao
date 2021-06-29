package alilabs

// AlibabaAilabUserProfileGetResult 
type AlibabaAilabUserProfileGetResult struct {

    // 数据model
    
    Result   *BasicUserInfo `json:"result,omitempty" xml:"result,omitempty"`
    

    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 200 成功，其他 失败
    
    StatusCode   int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
    

}
