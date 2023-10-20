package damai

import (
	"sync"
)

// TicketPaperFormatIdOpenParam 结构体
type TicketPaperFormatIdOpenParam struct {
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// 票纸版式id
	TicketPaperFormatId int64 `json:"ticket_paper_format_id,omitempty" xml:"ticket_paper_format_id,omitempty"`
}

var poolTicketPaperFormatIdOpenParam = sync.Pool{
	New: func() any {
		return new(TicketPaperFormatIdOpenParam)
	},
}

// GetTicketPaperFormatIdOpenParam() 从对象池中获取TicketPaperFormatIdOpenParam
func GetTicketPaperFormatIdOpenParam() *TicketPaperFormatIdOpenParam {
	return poolTicketPaperFormatIdOpenParam.Get().(*TicketPaperFormatIdOpenParam)
}

// ReleaseTicketPaperFormatIdOpenParam 释放TicketPaperFormatIdOpenParam
func ReleaseTicketPaperFormatIdOpenParam(v *TicketPaperFormatIdOpenParam) {
	v.SupplierSecret = ""
	v.SystemId = 0
	v.TicketPaperFormatId = 0
	poolTicketPaperFormatIdOpenParam.Put(v)
}
