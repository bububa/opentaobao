package alihouse

// TradeMerchantProofDto 结构体
type TradeMerchantProofDto struct {
	// 文件URL
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 文件类型
	FileType int64 `json:"file_type,omitempty" xml:"file_type,omitempty"`
}
