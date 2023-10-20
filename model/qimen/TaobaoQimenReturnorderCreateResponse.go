package qimen

// TaobaoqimenreturnordercreateResponse 结构体
type TaobaoqimenreturnordercreateResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 仓储系统退货单编码
	ReturnOrderId string `json:"returnOrderId,omitempty" xml:"returnOrderId,omitempty"`
}
