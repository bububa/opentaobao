package waybill

// CmdRenderParams 结构体
type CmdRenderParams struct {
	// 需要打印的文档，包括模板地址、打印数据
	Document *RenderDocument `json:"document,omitempty" xml:"document,omitempty"`
	// 打印机名称
	PrinterName string `json:"printer_name,omitempty" xml:"printer_name,omitempty"`
	// 客户端ID
	ClientId string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// 客户端类型：weixin/alipay/native
	ClientType string `json:"client_type,omitempty" xml:"client_type,omitempty"`
	// 打印配置
	Config *RenderConfig `json:"config,omitempty" xml:"config,omitempty"`
}
