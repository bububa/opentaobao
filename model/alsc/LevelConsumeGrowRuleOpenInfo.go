package alsc

import (
	"sync"
)

// LevelConsumeGrowRuleOpenInfo 结构体
type LevelConsumeGrowRuleOpenInfo struct {
	// 创建人
	CreateBy string `json:"create_by,omitempty" xml:"create_by,omitempty"`
	// 扩展字段
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 等级ID
	LevelId string `json:"level_id,omitempty" xml:"level_id,omitempty"`
	// 等级名称
	LevelName string `json:"level_name,omitempty" xml:"level_name,omitempty"`
	// 等级编码
	LevelNo string `json:"level_no,omitempty" xml:"level_no,omitempty"`
	// 更新人
	UpdateBy string `json:"update_by,omitempty" xml:"update_by,omitempty"`
	// 每消费金额，单位：分
	PerConsume int64 `json:"per_consume,omitempty" xml:"per_consume,omitempty"`
	// 每消费金额对应可获得的成长值
	PerGrowth int64 `json:"per_growth,omitempty" xml:"per_growth,omitempty"`
	// 是否已经删除
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
}

var poolLevelConsumeGrowRuleOpenInfo = sync.Pool{
	New: func() any {
		return new(LevelConsumeGrowRuleOpenInfo)
	},
}

// GetLevelConsumeGrowRuleOpenInfo() 从对象池中获取LevelConsumeGrowRuleOpenInfo
func GetLevelConsumeGrowRuleOpenInfo() *LevelConsumeGrowRuleOpenInfo {
	return poolLevelConsumeGrowRuleOpenInfo.Get().(*LevelConsumeGrowRuleOpenInfo)
}

// ReleaseLevelConsumeGrowRuleOpenInfo 释放LevelConsumeGrowRuleOpenInfo
func ReleaseLevelConsumeGrowRuleOpenInfo(v *LevelConsumeGrowRuleOpenInfo) {
	v.CreateBy = ""
	v.ExtInfo = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.LevelId = ""
	v.LevelName = ""
	v.LevelNo = ""
	v.UpdateBy = ""
	v.PerConsume = 0
	v.PerGrowth = 0
	v.Deleted = false
	poolLevelConsumeGrowRuleOpenInfo.Put(v)
}
