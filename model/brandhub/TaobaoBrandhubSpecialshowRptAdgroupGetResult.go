package brandhub

import (
	"sync"
)

// TaobaoBrandhubSpecialshowRptAdgroupGetResult 结构体
type TaobaoBrandhubSpecialshowRptAdgroupGetResult struct {
	// 日期
	Thedate string `json:"thedate,omitempty" xml:"thedate,omitempty"`
	// 计划名称
	SolutionName string `json:"solution_name,omitempty" xml:"solution_name,omitempty"`
	// 定向条件
	TargetName string `json:"target_name,omitempty" xml:"target_name,omitempty"`
	// 单元名称
	TaskName string `json:"task_name,omitempty" xml:"task_name,omitempty"`
	// 展现
	Impression string `json:"impression,omitempty" xml:"impression,omitempty"`
	// 点击
	Click string `json:"click,omitempty" xml:"click,omitempty"`
	// 访客数
	Uv string `json:"uv,omitempty" xml:"uv,omitempty"`
	// 点击访客数
	ClickUv string `json:"click_uv,omitempty" xml:"click_uv,omitempty"`
	// 点击率
	Ctr string `json:"ctr,omitempty" xml:"ctr,omitempty"`
	// 访客点击率
	UvCtr string `json:"uv_ctr,omitempty" xml:"uv_ctr,omitempty"`
	// 计划id
	SolutionId int64 `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
	// 单元id
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

var poolTaobaoBrandhubSpecialshowRptAdgroupGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoBrandhubSpecialshowRptAdgroupGetResult)
	},
}

// GetTaobaoBrandhubSpecialshowRptAdgroupGetResult() 从对象池中获取TaobaoBrandhubSpecialshowRptAdgroupGetResult
func GetTaobaoBrandhubSpecialshowRptAdgroupGetResult() *TaobaoBrandhubSpecialshowRptAdgroupGetResult {
	return poolTaobaoBrandhubSpecialshowRptAdgroupGetResult.Get().(*TaobaoBrandhubSpecialshowRptAdgroupGetResult)
}

// ReleaseTaobaoBrandhubSpecialshowRptAdgroupGetResult 释放TaobaoBrandhubSpecialshowRptAdgroupGetResult
func ReleaseTaobaoBrandhubSpecialshowRptAdgroupGetResult(v *TaobaoBrandhubSpecialshowRptAdgroupGetResult) {
	v.Thedate = ""
	v.SolutionName = ""
	v.TargetName = ""
	v.TaskName = ""
	v.Impression = ""
	v.Click = ""
	v.Uv = ""
	v.ClickUv = ""
	v.Ctr = ""
	v.UvCtr = ""
	v.SolutionId = 0
	v.TaskId = 0
	poolTaobaoBrandhubSpecialshowRptAdgroupGetResult.Put(v)
}
