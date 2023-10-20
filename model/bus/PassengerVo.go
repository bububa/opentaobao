package bus

import (
	"sync"
)

// PassengerVo 结构体
type PassengerVo struct {
	// 乘客证件号码
	RiderCertNumber string `json:"rider_cert_number,omitempty" xml:"rider_cert_number,omitempty"`
	// 乘客证件类型
	RiderCertType string `json:"rider_cert_type,omitempty" xml:"rider_cert_type,omitempty"`
	// 乘客姓名
	RiderName string `json:"rider_name,omitempty" xml:"rider_name,omitempty"`
	// 服务费
	ServiceCharge int64 `json:"service_charge,omitempty" xml:"service_charge,omitempty"`
	// 票价
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
}

var poolPassengerVo = sync.Pool{
	New: func() any {
		return new(PassengerVo)
	},
}

// GetPassengerVo() 从对象池中获取PassengerVo
func GetPassengerVo() *PassengerVo {
	return poolPassengerVo.Get().(*PassengerVo)
}

// ReleasePassengerVo 释放PassengerVo
func ReleasePassengerVo(v *PassengerVo) {
	v.RiderCertNumber = ""
	v.RiderCertType = ""
	v.RiderName = ""
	v.ServiceCharge = 0
	v.TicketPrice = 0
	poolPassengerVo.Put(v)
}
