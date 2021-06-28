package alihealth

// ServiceResult 
/* model for simplify = false
type ServiceResult struct {

    // errMessage
    
    ErrMessage   string `json:"err_message,omitempty"`
    

    // token
    
    Data   string `json:"data,omitempty"`
    

    // errCode
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// ServiceResult 
type ServiceResult struct {

    // errMessage
    ErrMessage   string `json:"err_message,omitempty"`

    // token
    Data   string `json:"data,omitempty"`

    // errCode
    ErrCode   string `json:"err_code,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
