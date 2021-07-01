package security

// RpUploadResult 结构体
type RpUploadResult struct {
	// uploadId
	UploadId string `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
	// uploadStatus
	UploadStatus *RpUploadStatus `json:"upload_status,omitempty" xml:"upload_status,omitempty"`
}
