package security

// RpUploadResult 
type RpUploadResult struct {

    // uploadId
    UploadId   string `json:"upload_id,omitempty"`

    // uploadStatus
    UploadStatus   *RpUploadStatus `json:"upload_status,omitempty"`

}
