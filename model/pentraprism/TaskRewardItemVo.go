package pentraprism

import (
	"sync"
)

// TaskRewardItemVo 结构体
type TaskRewardItemVo struct {
	// 奖励发放图标
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 激励模式 用于展示 &#34;NONE&#34;表示无激励
	Mode string `json:"mode,omitempty" xml:"mode,omitempty"`
	// 总库存格式化带单位 万/亿
	OwnCountText string `json:"own_count_text,omitempty" xml:"own_count_text,omitempty"`
	// 奖励类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 单位 用于展示
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 基础发放量
	BaseCount int64 `json:"base_count,omitempty" xml:"base_count,omitempty"`
	// 奖励模板ID
	ConfigId int64 `json:"config_id,omitempty" xml:"config_id,omitempty"`
	// 奖励激励值
	Encourage int64 `json:"encourage,omitempty" xml:"encourage,omitempty"`
	// 最终奖励发放数量
	FinalCount int64 `json:"final_count,omitempty" xml:"final_count,omitempty"`
	// 奖励区间最大值 用于展示
	MaxCount int64 `json:"max_count,omitempty" xml:"max_count,omitempty"`
	// 奖励区间最小值 用于展示
	MinCount int64 `json:"min_count,omitempty" xml:"min_count,omitempty"`
	// 库存量
	OwnCount int64 `json:"own_count,omitempty" xml:"own_count,omitempty"`
	// 成就点
	PointCount int64 `json:"point_count,omitempty" xml:"point_count,omitempty"`
}

var poolTaskRewardItemVo = sync.Pool{
	New: func() any {
		return new(TaskRewardItemVo)
	},
}

// GetTaskRewardItemVo() 从对象池中获取TaskRewardItemVo
func GetTaskRewardItemVo() *TaskRewardItemVo {
	return poolTaskRewardItemVo.Get().(*TaskRewardItemVo)
}

// ReleaseTaskRewardItemVo 释放TaskRewardItemVo
func ReleaseTaskRewardItemVo(v *TaskRewardItemVo) {
	v.Icon = ""
	v.Mode = ""
	v.OwnCountText = ""
	v.Type = ""
	v.Unit = ""
	v.BaseCount = 0
	v.ConfigId = 0
	v.Encourage = 0
	v.FinalCount = 0
	v.MaxCount = 0
	v.MinCount = 0
	v.OwnCount = 0
	v.PointCount = 0
	poolTaskRewardItemVo.Put(v)
}
