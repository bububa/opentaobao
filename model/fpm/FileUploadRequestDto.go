package fpm

// FileUploadRequestDto 结构体
type FileUploadRequestDto struct {
	// 应用代码(必填)
	AppCode string `json:"app_code,omitempty" xml:"app_code,omitempty"`
	// 签名字符串
	OuterSystemSignStr string `json:"outer_system_sign_str,omitempty" xml:"outer_system_sign_str,omitempty"`
	// 加密字符串
	OuterSystemEncryptStr string `json:"outer_system_encrypt_str,omitempty" xml:"outer_system_encrypt_str,omitempty"`
}
