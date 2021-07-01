package damai

// ThirdPaperFormatPushOpenParam 结构体
type ThirdPaperFormatPushOpenParam struct {
	// 票纸版式id
	PaperFormatId int64 `json:"paper_format_id,omitempty" xml:"paper_format_id,omitempty"`
	// 票纸版式名称
	PaperFormatName string `json:"paper_format_name,omitempty" xml:"paper_format_name,omitempty"`
	// 打印类型
	PrintType int64 `json:"print_type,omitempty" xml:"print_type,omitempty"`
	// 推送时间
	PushTime string `json:"push_time,omitempty" xml:"push_time,omitempty"`
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
}
