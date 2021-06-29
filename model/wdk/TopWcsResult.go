package wdk

// TopWcsResult 
type TopWcsResult struct {
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // errorCode
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // errorCode
    ServiceErrorCode   string `json:"service_error_code,omitempty" xml:"service_error_code,omitempty"`
    // errorMsg
    ServiceErrorMsg   string `json:"service_error_msg,omitempty" xml:"service_error_msg,omitempty"`
    // success
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // list
    List   []WcsContainerAssignedToConveyorDTO `json:"list,omitempty" xml:"list>wcs_container_assigned_to_conveyor_dto,omitempty"`
}
