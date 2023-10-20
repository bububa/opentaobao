package ieagency

import (
	"sync"
)

// IePassengerTicketVo 结构体
type IePassengerTicketVo struct {
	// 票号
	TicketNos []string `json:"ticket_nos,omitempty" xml:"ticket_nos>string,omitempty"`
	// 证件号码
	CertNo string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// 乘机人姓名（英文）
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 代理人pnr
	Cpnr string `json:"cpnr,omitempty" xml:"cpnr,omitempty"`
	// 航空公司pnr
	Bpnr string `json:"bpnr,omitempty" xml:"bpnr,omitempty"`
	// 出票渠道，可以填1E、1A、1S、1G、官网、其他
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
}

var poolIePassengerTicketVo = sync.Pool{
	New: func() any {
		return new(IePassengerTicketVo)
	},
}

// GetIePassengerTicketVo() 从对象池中获取IePassengerTicketVo
func GetIePassengerTicketVo() *IePassengerTicketVo {
	return poolIePassengerTicketVo.Get().(*IePassengerTicketVo)
}

// ReleaseIePassengerTicketVo 释放IePassengerTicketVo
func ReleaseIePassengerTicketVo(v *IePassengerTicketVo) {
	v.TicketNos = v.TicketNos[:0]
	v.CertNo = ""
	v.PassengerName = ""
	v.Cpnr = ""
	v.Bpnr = ""
	v.Channel = ""
	poolIePassengerTicketVo.Put(v)
}
