package fans

import (
	"sync"
)

// CreateCashPoolParamDo 结构体
type CreateCashPoolParamDo struct {
	// 红包使用开始时间
	UseEndTime string `json:"use_end_time,omitempty" xml:"use_end_time,omitempty"`
	// 红包使用结束时间
	UseStartTime string `json:"use_start_time,omitempty" xml:"use_start_time,omitempty"`
	// 奖金池标题
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 奖金池描述
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 活动id
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 活动url
	ActivityUrl string `json:"activity_url,omitempty" xml:"activity_url,omitempty"`
	// 开奖时间
	DrawTime string `json:"draw_time,omitempty" xml:"draw_time,omitempty"`
	// 活动开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 活动结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 开始答题时间
	BeginQaTime string `json:"begin_qa_time,omitempty" xml:"begin_qa_time,omitempty"`
	// 奖金池总额度
	CashValue int64 `json:"cash_value,omitempty" xml:"cash_value,omitempty"`
}

var poolCreateCashPoolParamDo = sync.Pool{
	New: func() any {
		return new(CreateCashPoolParamDo)
	},
}

// GetCreateCashPoolParamDo() 从对象池中获取CreateCashPoolParamDo
func GetCreateCashPoolParamDo() *CreateCashPoolParamDo {
	return poolCreateCashPoolParamDo.Get().(*CreateCashPoolParamDo)
}

// ReleaseCreateCashPoolParamDo 释放CreateCashPoolParamDo
func ReleaseCreateCashPoolParamDo(v *CreateCashPoolParamDo) {
	v.UseEndTime = ""
	v.UseStartTime = ""
	v.Desc = ""
	v.Title = ""
	v.ActivityId = ""
	v.ActivityUrl = ""
	v.DrawTime = ""
	v.StartTime = ""
	v.EndTime = ""
	v.BeginQaTime = ""
	v.CashValue = 0
	poolCreateCashPoolParamDo.Put(v)
}
