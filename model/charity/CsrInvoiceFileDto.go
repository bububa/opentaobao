package charity

// CsrInvoiceFileDto 结构体
type CsrInvoiceFileDto struct {
	// 纸质发票必填-快递单号
	ExpressNo string `json:"express_no,omitempty" xml:"express_no,omitempty"`
	// 纸质发票必填-快递服务商
	ExpressType string `json:"express_type,omitempty" xml:"express_type,omitempty"`
	// 电子发票必填-发票文件地址
	FileUrl string `json:"file_url,omitempty" xml:"file_url,omitempty"`
	// 电子发票必填-发票文件名称
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
}
