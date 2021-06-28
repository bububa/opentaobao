package wdk

// TaobaoWdkEquipmentConveyorConveyorinfoGetResult 
/* model for simplify = false
type TaobaoWdkEquipmentConveyorConveyorinfoGetResult struct {

    // errorCode
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

    // model
    
    Model  *struct {
        WcsConveyorInfoDto  *WcsConveyorInfoDto `json:"wcs_conveyor_info_dto,omitempty"`
    } `json:"model,omitempty"`
    

    // errorMsg
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

}
*/

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
