package tmallsc

// TmallWorkcardIdentifyResult 
type TmallWorkcardIdentifyResult struct {

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

    // errorCode
    ErrorCode   int64 `json:"error_code,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // 返回void
    Value   string `json:"value,omitempty"`

    // gmtModified
    GmtModified   string `json:"gmt_modified,omitempty"`

    // gmtCreate
    GmtCreate   string `json:"gmt_create,omitempty"`

}
