package media

// GenerateTokenRequest 结构体
type GenerateTokenRequest struct {
	// 请求策略
	UploadPolicy *UploadPolicy `json:"upload_policy,omitempty" xml:"upload_policy,omitempty"`
}
