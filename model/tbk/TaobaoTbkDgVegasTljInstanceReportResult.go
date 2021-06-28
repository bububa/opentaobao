package tbk

// TaobaoTbkDgVegasTljInstanceReportResult 
/* model for simplify = false
type TaobaoTbkDgVegasTljInstanceReportResult struct {

    // model
    
    Model  *struct {
        TljInstanceReportDto  *TljInstanceReportDto `json:"tlj_instance_report_dto,omitempty"`
    } `json:"model,omitempty"`
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TaobaoTbkDgVegasTljInstanceReportResult 
type TaobaoTbkDgVegasTljInstanceReportResult struct {

    // model
    Model   *TljInstanceReportDto `json:"model,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
