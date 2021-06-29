package sungari

// TaobaoCloudbridgeCaseinvestExecuteResult 
type TaobaoCloudbridgeCaseinvestExecuteResult struct {

    // 接口调用是否成功，1：成功；2：失败
    
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    

    // data值，JSON数据，可转换成对应的结果
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    

    // 说明
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

}
