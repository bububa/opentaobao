package trade

// TmallAscpOrdersSaleCreateResultDO 
type TmallAscpOrdersSaleCreateResultDO struct {
    // errorMessage
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // module
    Module   string `json:"module,omitempty" xml:"module,omitempty"`
    // totalCount
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    // errorCode
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
