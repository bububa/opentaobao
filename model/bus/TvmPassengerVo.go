package bus

import (
	"sync"
)

// TvmPassengerVo 结构体
type TvmPassengerVo struct {
	// 保险信息
	TvmInsuranceInfos []TvmInsuranceInfo `json:"tvm_insurance_infos,omitempty" xml:"tvm_insurance_infos>tvm_insurance_info,omitempty"`
	// 电子票内容
	AgentEticket string `json:"agent_eticket,omitempty" xml:"agent_eticket,omitempty"`
	// 票号ID
	AgentTicketId string `json:"agent_ticket_id,omitempty" xml:"agent_ticket_id,omitempty"`
	// 身份证号
	RiderCertNumber string `json:"rider_cert_number,omitempty" xml:"rider_cert_number,omitempty"`
	// 证件类型
	RiderCertType string `json:"rider_cert_type,omitempty" xml:"rider_cert_type,omitempty"`
	// 姓名
	RiderName string `json:"rider_name,omitempty" xml:"rider_name,omitempty"`
	// 座位号
	SeatNumber string `json:"seat_number,omitempty" xml:"seat_number,omitempty"`
	// 标准票价
	FullPrice int64 `json:"full_price,omitempty" xml:"full_price,omitempty"`
	// 服务费
	ServiceCharge int64 `json:"service_charge,omitempty" xml:"service_charge,omitempty"`
	// 实际票价
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// 保险费 (单位分)
	InsurePrice int64 `json:"insure_price,omitempty" xml:"insure_price,omitempty"`
	// 乘客类型：0：成人，1：儿童，2：携童
	TicketType int64 `json:"ticket_type,omitempty" xml:"ticket_type,omitempty"`
	// 折扣类型：1：特殊证件，2：特殊线路，3其他。如果是折扣票必须填写
	DiscountType int64 `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
	// 是否带有儿童
	HasChildren bool `json:"has_children,omitempty" xml:"has_children,omitempty"`
}

var poolTvmPassengerVo = sync.Pool{
	New: func() any {
		return new(TvmPassengerVo)
	},
}

// GetTvmPassengerVo() 从对象池中获取TvmPassengerVo
func GetTvmPassengerVo() *TvmPassengerVo {
	return poolTvmPassengerVo.Get().(*TvmPassengerVo)
}

// ReleaseTvmPassengerVo 释放TvmPassengerVo
func ReleaseTvmPassengerVo(v *TvmPassengerVo) {
	v.TvmInsuranceInfos = v.TvmInsuranceInfos[:0]
	v.AgentEticket = ""
	v.AgentTicketId = ""
	v.RiderCertNumber = ""
	v.RiderCertType = ""
	v.RiderName = ""
	v.SeatNumber = ""
	v.FullPrice = 0
	v.ServiceCharge = 0
	v.TicketPrice = 0
	v.InsurePrice = 0
	v.TicketType = 0
	v.DiscountType = 0
	v.HasChildren = false
	poolTvmPassengerVo.Put(v)
}
