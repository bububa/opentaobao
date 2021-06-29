package eleenterpriseemployee

// EnterpriseData 
type EnterpriseData struct {

    // 成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 错误信息
    
    ErrorMsgs   []ErrorMsg `json:"error_msgs,omitempty" xml:"error_msgs,omitempty"`
    

}
