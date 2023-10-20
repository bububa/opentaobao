package iotticket

import (
	"sync"
)

// AcceptTicketV2TopRequest 结构体
type AcceptTicketV2TopRequest struct {
	// 维修方案 depot_repair:寄回维修;parts_replacement:配件更换;onsite_repair:上门维修;remote_solution:远程解决;transfer_to_customer_service:转单给菜鸟
	MaintenanceModeCode string `json:"maintenance_mode_code,omitempty" xml:"maintenance_mode_code,omitempty"`
	// 上门人员Id
	OnsiteStaffId string `json:"onsite_staff_id,omitempty" xml:"onsite_staff_id,omitempty"`
	// 收件人名称
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 上门地址
	OnsiteAddress string `json:"onsite_address,omitempty" xml:"onsite_address,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 操作人手机号
	OperatorPhone string `json:"operator_phone,omitempty" xml:"operator_phone,omitempty"`
	// 操作人名称
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 上门人员名称
	OnsiteStaffName string `json:"onsite_staff_name,omitempty" xml:"onsite_staff_name,omitempty"`
	// 上门时间
	OnsiteTime string `json:"onsite_time,omitempty" xml:"onsite_time,omitempty"`
	// 服务商编码
	SpCode string `json:"sp_code,omitempty" xml:"sp_code,omitempty"`
	// 收件地址
	ReceiverAddress string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
	// 上门人联系方式
	OnsiteStaffPhone string `json:"onsite_staff_phone,omitempty" xml:"onsite_staff_phone,omitempty"`
	// 收件人联系方式
	ReceiverPhone string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
	// 扩展字段
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 远程解决方式 1:不想修了;2:已经修好了;3:费用太贵;4:其他原因
	SolutionRemark string `json:"solution_remark,omitempty" xml:"solution_remark,omitempty"`
	// 操作人Id
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 工单Id
	TicketId int64 `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
}

var poolAcceptTicketV2TopRequest = sync.Pool{
	New: func() any {
		return new(AcceptTicketV2TopRequest)
	},
}

// GetAcceptTicketV2TopRequest() 从对象池中获取AcceptTicketV2TopRequest
func GetAcceptTicketV2TopRequest() *AcceptTicketV2TopRequest {
	return poolAcceptTicketV2TopRequest.Get().(*AcceptTicketV2TopRequest)
}

// ReleaseAcceptTicketV2TopRequest 释放AcceptTicketV2TopRequest
func ReleaseAcceptTicketV2TopRequest(v *AcceptTicketV2TopRequest) {
	v.MaintenanceModeCode = ""
	v.OnsiteStaffId = ""
	v.ReceiverName = ""
	v.OnsiteAddress = ""
	v.Remark = ""
	v.OperatorPhone = ""
	v.OperatorName = ""
	v.OnsiteStaffName = ""
	v.OnsiteTime = ""
	v.SpCode = ""
	v.ReceiverAddress = ""
	v.OnsiteStaffPhone = ""
	v.ReceiverPhone = ""
	v.Feature = ""
	v.SolutionRemark = ""
	v.OperatorId = ""
	v.TicketId = 0
	poolAcceptTicketV2TopRequest.Put(v)
}
