package ship

import (
	"sync"
)

// ShipAgentConfirmBookPassengerInfo 结构体
type ShipAgentConfirmBookPassengerInfo struct {
	// 票信息
	TicketList []ShipAgentConfirmBookTicketInfo `json:"ticket_list,omitempty" xml:"ticket_list>ship_agent_confirm_book_ticket_info,omitempty"`
	// 乘客证件号
	PassengerCertNo string `json:"passenger_cert_no,omitempty" xml:"passenger_cert_no,omitempty"`
	// 乘客证件类型
	PassengerCertType string `json:"passenger_cert_type,omitempty" xml:"passenger_cert_type,omitempty"`
	// 乘客id
	PassengerId string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// 乘客姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
}

var poolShipAgentConfirmBookPassengerInfo = sync.Pool{
	New: func() any {
		return new(ShipAgentConfirmBookPassengerInfo)
	},
}

// GetShipAgentConfirmBookPassengerInfo() 从对象池中获取ShipAgentConfirmBookPassengerInfo
func GetShipAgentConfirmBookPassengerInfo() *ShipAgentConfirmBookPassengerInfo {
	return poolShipAgentConfirmBookPassengerInfo.Get().(*ShipAgentConfirmBookPassengerInfo)
}

// ReleaseShipAgentConfirmBookPassengerInfo 释放ShipAgentConfirmBookPassengerInfo
func ReleaseShipAgentConfirmBookPassengerInfo(v *ShipAgentConfirmBookPassengerInfo) {
	v.TicketList = v.TicketList[:0]
	v.PassengerCertNo = ""
	v.PassengerCertType = ""
	v.PassengerId = ""
	v.PassengerName = ""
	poolShipAgentConfirmBookPassengerInfo.Put(v)
}
