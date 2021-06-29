package wdk

// TaobaoWdkEquipmentConveyorContainerinfoGetResult 
type TaobaoWdkEquipmentConveyorContainerinfoGetResult struct {
    // result
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
    // errorDesc
    ErrorDesc   string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
    // errorCode
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // totelNum
    TotelNum   int64 `json:"totel_num,omitempty" xml:"totel_num,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
