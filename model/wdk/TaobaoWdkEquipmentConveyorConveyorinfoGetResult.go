package wdk

// TaobaoWdkEquipmentConveyorConveyorinfoGetResult 
type TaobaoWdkEquipmentConveyorConveyorinfoGetResult struct {

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // model
    Model   *WcsConveyorInfoDto `json:"model,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

}
