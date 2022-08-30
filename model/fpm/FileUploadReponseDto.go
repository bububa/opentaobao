package fpm

// FileUploadReponseDto 结构体
type FileUploadReponseDto struct {
	// outerSystemSignStr
	OuterSystemSignStr string `json:"outer_system_sign_str,omitempty" xml:"outer_system_sign_str,omitempty"`
	// outerSystemEncryptStr
	OuterSystemEncryptStr string `json:"outer_system_encrypt_str,omitempty" xml:"outer_system_encrypt_str,omitempty"`
}
