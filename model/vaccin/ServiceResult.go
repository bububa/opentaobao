package vaccin

// ServiceResult 
type ServiceResult struct {

    // data
    
    DivisionList   []DivisionDto `json:"division_list,omitempty" xml:"division_list,omitempty"`
    

    // errMessage
    
    ErrMessage   string `json:"err_message,omitempty" xml:"err_message,omitempty"`
    

    // errCode
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
