package icbuassurance

// AssuranceAccountResult 
type AssuranceAccountResult struct {

    // errorMessage
    
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    

    // value
    
    Value   *AssuranceFlag `json:"value,omitempty" xml:"value,omitempty"`
    

    // errorCode
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
