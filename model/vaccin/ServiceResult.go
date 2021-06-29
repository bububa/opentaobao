package vaccin

// ServiceResult 
type ServiceResult struct {
    // data
    DivisionList   []DivisionDTO `json:"division_list,omitempty" xml:"division_list>division_dto,omitempty"`
    // errMessage
    ErrMessage   string `json:"err_message,omitempty" xml:"err_message,omitempty"`
    // errCode
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
