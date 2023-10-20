package aetask

import (
	"sync"
)

// Materials 结构体
type Materials struct {
	// 物料类型：shop/item/url
	MaterialUrlType string `json:"material_url_type,omitempty" xml:"material_url_type,omitempty"`
	// 物料url
	MaterialUrl string `json:"material_url,omitempty" xml:"material_url,omitempty"`
	// 未兑换记录
	UnIssueRecord string `json:"un_issue_record,omitempty" xml:"un_issue_record,omitempty"`
	// 扩展待用
	DetailIds string `json:"detail_ids,omitempty" xml:"detail_ids,omitempty"`
	// 幂等id
	IdempotentId string `json:"idempotent_id,omitempty" xml:"idempotent_id,omitempty"`
	// 未完成图标
	Icon4UnFinish string `json:"icon4_un_finish,omitempty" xml:"icon4_un_finish,omitempty"`
	// 扩展字段
	ExtendInfo string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
	// 完成图标
	Icon4Finished string `json:"icon4_finished,omitempty" xml:"icon4_finished,omitempty"`
	// 配置项
	BehaviorConfig string `json:"behavior_config,omitempty" xml:"behavior_config,omitempty"`
	// 算法打点标志
	Trace string `json:"trace,omitempty" xml:"trace,omitempty"`
	// 主标题
	MainTitle string `json:"main_title,omitempty" xml:"main_title,omitempty"`
	// 副标题
	SecondTitle string `json:"second_title,omitempty" xml:"second_title,omitempty"`
	// 图标
	IconUrl string `json:"icon_url,omitempty" xml:"icon_url,omitempty"`
	// 利益点
	InterestNum int64 `json:"interest_num,omitempty" xml:"interest_num,omitempty"`
	// 参与记录id
	TaskInstanceId int64 `json:"task_instance_id,omitempty" xml:"task_instance_id,omitempty"`
	// 任务类型：0浏览
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 任务状态
	InstanceStatus int64 `json:"instance_status,omitempty" xml:"instance_status,omitempty"`
	// 分组排序
	OrderGroup int64 `json:"order_group,omitempty" xml:"order_group,omitempty"`
	// 已参与次数
	TimesJoined int64 `json:"times_joined,omitempty" xml:"times_joined,omitempty"`
	// 物料id
	MaterialConfigId int64 `json:"material_config_id,omitempty" xml:"material_config_id,omitempty"`
	// 可玩次数
	TimesLimit int64 `json:"times_limit,omitempty" xml:"times_limit,omitempty"`
	// 任务id
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
	// 是否原地跳转
	BrowseZeroFlag bool `json:"browse_zero_flag,omitempty" xml:"browse_zero_flag,omitempty"`
}

var poolMaterials = sync.Pool{
	New: func() any {
		return new(Materials)
	},
}

// GetMaterials() 从对象池中获取Materials
func GetMaterials() *Materials {
	return poolMaterials.Get().(*Materials)
}

// ReleaseMaterials 释放Materials
func ReleaseMaterials(v *Materials) {
	v.MaterialUrlType = ""
	v.MaterialUrl = ""
	v.UnIssueRecord = ""
	v.DetailIds = ""
	v.IdempotentId = ""
	v.Icon4UnFinish = ""
	v.ExtendInfo = ""
	v.Icon4Finished = ""
	v.BehaviorConfig = ""
	v.Trace = ""
	v.MainTitle = ""
	v.SecondTitle = ""
	v.IconUrl = ""
	v.InterestNum = 0
	v.TaskInstanceId = 0
	v.Type = 0
	v.InstanceStatus = 0
	v.OrderGroup = 0
	v.TimesJoined = 0
	v.MaterialConfigId = 0
	v.TimesLimit = 0
	v.TaskId = 0
	v.BrowseZeroFlag = false
	poolMaterials.Put(v)
}
