package baoxian

// UploadResult 
/* model for simplify = false
type UploadResult struct {

    // model
    
    Model  *struct {
        InsAttachmentUploadVo  *InsAttachmentUploadVo `json:"ins_attachment_upload_vo,omitempty"`
    } `json:"model,omitempty"`
    

    // errorCode
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // errorMsg
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // isSuccess
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

// UploadResult 
type UploadResult struct {

    // model
    Model   *InsAttachmentUploadVo `json:"model,omitempty"`

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

    // isSuccess
    IsSuccess   bool `json:"is_success,omitempty"`

}
