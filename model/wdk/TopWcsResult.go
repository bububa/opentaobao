package wdk

// TopWcsResult 
/* model for simplify = false
type TopWcsResult struct {

    // success
    
    Success   bool `json:"success,omitempty"`
    

    // errorCode
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // errorMsg
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // errorCode
    
    ServiceErrorCode   string `json:"service_error_code,omitempty"`
    

    // errorMsg
    
    ServiceErrorMsg   string `json:"service_error_msg,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // list
    
    List  struct {
        WcsContainerAssignedToConveyorDto  []WcsContainerAssignedToConveyorDto `json:"wcs_container_assigned_to_conveyor_dto,omitempty"`
    } `json:"list,omitempty"`
    

}
*/

// TopWcsResult 
type TopWcsResult struct {

    // success
    Success   bool `json:"success,omitempty"`

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

    // errorCode
    ServiceErrorCode   string `json:"service_error_code,omitempty"`

    // errorMsg
    ServiceErrorMsg   string `json:"service_error_msg,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

    // list
    List   []WcsContainerAssignedToConveyorDto `json:"list,omitempty"`

}
