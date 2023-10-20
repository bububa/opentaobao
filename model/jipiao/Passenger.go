package jipiao

import (
	"sync"
)

// Passenger 结构体
type Passenger struct {
	// 乘客生日
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 乘客证件号码
	CertNum string `json:"cert_num,omitempty" xml:"cert_num,omitempty"`
	// 乘客姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 改签前的pnr
	Pnr string `json:"pnr,omitempty" xml:"pnr,omitempty"`
	// 改签后的票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 改签前的票号
	OldTicketNo string `json:"old_ticket_no,omitempty" xml:"old_ticket_no,omitempty"`
}

var poolPassenger = sync.Pool{
	New: func() any {
		return new(Passenger)
	},
}

// GetPassenger() 从对象池中获取Passenger
func GetPassenger() *Passenger {
	return poolPassenger.Get().(*Passenger)
}

// ReleasePassenger 释放Passenger
func ReleasePassenger(v *Passenger) {
	v.Birthday = ""
	v.CertNum = ""
	v.PassengerName = ""
	v.Pnr = ""
	v.TicketNo = ""
	v.OldTicketNo = ""
	poolPassenger.Put(v)
}
