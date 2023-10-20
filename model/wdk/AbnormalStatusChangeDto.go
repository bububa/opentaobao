package wdk

import (
	"sync"
)

// AbnormalStatusChangeDto 结构体
type AbnormalStatusChangeDto struct {
	// 是否终态
	IsFinal string `json:"is_final,omitempty" xml:"is_final,omitempty"`
	// 当前经纬度
	CurrentLngLat string `json:"current_lng_lat,omitempty" xml:"current_lng_lat,omitempty"`
	// 操作时间
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 处理人名称
	OperateUserName string `json:"operate_user_name,omitempty" xml:"operate_user_name,omitempty"`
	// 处理人
	OperateUserCode string `json:"operate_user_code,omitempty" xml:"operate_user_code,omitempty"`
	// 节点Code
	NodeCode string `json:"node_code,omitempty" xml:"node_code,omitempty"`
	// 节点类型
	NodeType string `json:"node_type,omitempty" xml:"node_type,omitempty"`
	// 异常受理单ID
	AbnormalAcceptId string `json:"abnormal_accept_id,omitempty" xml:"abnormal_accept_id,omitempty"`
	// 异常单ID
	AbnormalOrderId string `json:"abnormal_order_id,omitempty" xml:"abnormal_order_id,omitempty"`
	// 状态变更类型
	StatusChangeType string `json:"status_change_type,omitempty" xml:"status_change_type,omitempty"`
	// 异常协同单ID
	AbnormalCoordinationId string `json:"abnormal_coordination_id,omitempty" xml:"abnormal_coordination_id,omitempty"`
	// 参数
	AbnormalContentDto *AbnormalContentDto `json:"abnormal_content_dto,omitempty" xml:"abnormal_content_dto,omitempty"`
}

var poolAbnormalStatusChangeDto = sync.Pool{
	New: func() any {
		return new(AbnormalStatusChangeDto)
	},
}

// GetAbnormalStatusChangeDto() 从对象池中获取AbnormalStatusChangeDto
func GetAbnormalStatusChangeDto() *AbnormalStatusChangeDto {
	return poolAbnormalStatusChangeDto.Get().(*AbnormalStatusChangeDto)
}

// ReleaseAbnormalStatusChangeDto 释放AbnormalStatusChangeDto
func ReleaseAbnormalStatusChangeDto(v *AbnormalStatusChangeDto) {
	v.IsFinal = ""
	v.CurrentLngLat = ""
	v.OperateTime = ""
	v.OperateUserName = ""
	v.OperateUserCode = ""
	v.NodeCode = ""
	v.NodeType = ""
	v.AbnormalAcceptId = ""
	v.AbnormalOrderId = ""
	v.StatusChangeType = ""
	v.AbnormalCoordinationId = ""
	v.AbnormalContentDto = nil
	poolAbnormalStatusChangeDto.Put(v)
}
