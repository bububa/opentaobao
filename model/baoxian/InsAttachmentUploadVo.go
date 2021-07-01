package baoxian

// InsAttachmentUploadVo 结构体
type InsAttachmentUploadVo struct {
	// ossPath
	OssPath string `json:"oss_path,omitempty" xml:"oss_path,omitempty"`
	// eTag
	ETag string `json:"e_tag,omitempty" xml:"e_tag,omitempty"`
	// size
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
}
