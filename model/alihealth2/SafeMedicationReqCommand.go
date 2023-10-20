package alihealth2

import (
	"sync"
)

// SafeMedicationReqCommand 结构体
type SafeMedicationReqCommand struct {
	// 处方项，包括药品、用法用量以及购买量
	PrescriptionItems []PrescriptionItem `json:"prescription_items,omitempty" xml:"prescription_items>prescription_item,omitempty"`
	// 诊断信息，启用适应症和禁忌症规则下必传
	Diags []Diag `json:"diags,omitempty" xml:"diags>diag,omitempty"`
	// 关联订单
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 场景名
	SceneName string `json:"scene_name,omitempty" xml:"scene_name,omitempty"`
	// 租户名
	TenantCode string `json:"tenant_code,omitempty" xml:"tenant_code,omitempty"`
	// 患者，启用特殊人群规则情况下必传
	Patient *Patient `json:"patient,omitempty" xml:"patient,omitempty"`
}

var poolSafeMedicationReqCommand = sync.Pool{
	New: func() any {
		return new(SafeMedicationReqCommand)
	},
}

// GetSafeMedicationReqCommand() 从对象池中获取SafeMedicationReqCommand
func GetSafeMedicationReqCommand() *SafeMedicationReqCommand {
	return poolSafeMedicationReqCommand.Get().(*SafeMedicationReqCommand)
}

// ReleaseSafeMedicationReqCommand 释放SafeMedicationReqCommand
func ReleaseSafeMedicationReqCommand(v *SafeMedicationReqCommand) {
	v.PrescriptionItems = v.PrescriptionItems[:0]
	v.Diags = v.Diags[:0]
	v.OrderId = ""
	v.SceneName = ""
	v.TenantCode = ""
	v.Patient = nil
	poolSafeMedicationReqCommand.Put(v)
}
