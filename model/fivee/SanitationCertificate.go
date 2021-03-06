package fivee

// SanitationCertificate 结构体
type SanitationCertificate struct {
	// 批次信息
	BatchProducts []BatchProduct `json:"batch_products,omitempty" xml:"batch_products>batch_product,omitempty"`
	// 下载地址列表
	Urls []string `json:"urls,omitempty" xml:"urls>string,omitempty"`
	// 编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}
