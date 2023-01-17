package traveltrade

// NormalVisaETAInfo 结构体
type NormalVisaETAInfo struct {
	// 必填，电子签pdf文件名称。具体的pdf文件字节流信息请设置到父级参数的 fileBytes字段！！！
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
}
