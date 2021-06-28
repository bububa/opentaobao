package security

// RpUploadResult 
/* model for simplify = false
type RpUploadResult struct {

    // uploadId
    
    UploadId   string `json:"upload_id,omitempty"`
    

    // uploadStatus
    
    UploadStatus  *struct {
        RpUploadStatus  *RpUploadStatus `json:"rp_upload_status,omitempty"`
    } `json:"upload_status,omitempty"`
    

}
*/

// RpUploadResult 
type RpUploadResult struct {

    // uploadId
    UploadId   string `json:"upload_id,omitempty"`

    // uploadStatus
    UploadStatus   *RpUploadStatus `json:"upload_status,omitempty"`

}
