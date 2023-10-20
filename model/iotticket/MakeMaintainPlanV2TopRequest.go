package iotticket

import (
	"sync"
)

// MakeMaintainPlanV2TopRequest 结构体
type MakeMaintainPlanV2TopRequest struct {
	// 维修项
	IotMaintainPlanItemList []IotMaintainPlanItemTopRequest `json:"iot_maintain_plan_item_list,omitempty" xml:"iot_maintain_plan_item_list>iot_maintain_plan_item_top_request,omitempty"`
	// 问题列表（需要映射）
	ProblemTypeList []string `json:"problem_type_list,omitempty" xml:"problem_type_list>string,omitempty"`
	// 问题原因（需要映射）
	ProblemCauseList []string `json:"problem_cause_list,omitempty" xml:"problem_cause_list>string,omitempty"`
	// 事件类型（需要映射）
	EventTypeList []string `json:"event_type_list,omitempty" xml:"event_type_list>string,omitempty"`
	// 其它费用
	OtherFee string `json:"other_fee,omitempty" xml:"other_fee,omitempty"`
	// 操作人联系方式
	OperatorPhone string `json:"operator_phone,omitempty" xml:"operator_phone,omitempty"`
	// 操作人姓名
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 服务商编码
	SpCode string `json:"sp_code,omitempty" xml:"sp_code,omitempty"`
	// 扩展属性
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 费用描述
	FeeRemark string `json:"fee_remark,omitempty" xml:"fee_remark,omitempty"`
	// 保内保外（需要映射）
	WarrantyType string `json:"warranty_type,omitempty" xml:"warranty_type,omitempty"`
	// 支付图片二维码
	PayPictureUrl string `json:"pay_picture_url,omitempty" xml:"pay_picture_url,omitempty"`
	// 操作人Id
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 人工费用
	LaborExpense string `json:"labor_expense,omitempty" xml:"labor_expense,omitempty"`
	// 工单Id
	TicketId int64 `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
}

var poolMakeMaintainPlanV2TopRequest = sync.Pool{
	New: func() any {
		return new(MakeMaintainPlanV2TopRequest)
	},
}

// GetMakeMaintainPlanV2TopRequest() 从对象池中获取MakeMaintainPlanV2TopRequest
func GetMakeMaintainPlanV2TopRequest() *MakeMaintainPlanV2TopRequest {
	return poolMakeMaintainPlanV2TopRequest.Get().(*MakeMaintainPlanV2TopRequest)
}

// ReleaseMakeMaintainPlanV2TopRequest 释放MakeMaintainPlanV2TopRequest
func ReleaseMakeMaintainPlanV2TopRequest(v *MakeMaintainPlanV2TopRequest) {
	v.IotMaintainPlanItemList = v.IotMaintainPlanItemList[:0]
	v.ProblemTypeList = v.ProblemTypeList[:0]
	v.ProblemCauseList = v.ProblemCauseList[:0]
	v.EventTypeList = v.EventTypeList[:0]
	v.OtherFee = ""
	v.OperatorPhone = ""
	v.OperatorName = ""
	v.SpCode = ""
	v.Feature = ""
	v.FeeRemark = ""
	v.WarrantyType = ""
	v.PayPictureUrl = ""
	v.OperatorId = ""
	v.LaborExpense = ""
	v.TicketId = 0
	poolMakeMaintainPlanV2TopRequest.Put(v)
}
