package btrip

// BtmsResult 
type BtmsResult struct {

    // 分页结果对象。
    
    Module   *PagingResult `json:"module,omitempty" xml:"module,omitempty"`
    

    // 结果码；0：成功，非0：失败。
    
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
    

    // 结果描述。
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
    

    // 请求是否成功。
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
