package eleenterpriseemployee

// ErrorMsg 
type ErrorMsg struct {

    // 请求报文
    
    ReqBody   string `json:"req_body,omitempty" xml:"req_body,omitempty"`
    

    // 失败原因
    
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
    

}
