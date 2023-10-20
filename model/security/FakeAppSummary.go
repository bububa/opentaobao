package security

import (
	"sync"
)

// FakeAppSummary 结构体
type FakeAppSummary struct {
	// 仿冒应用总数(任务完成时才返回)
	FakeAppCount int64 `json:"fake_app_count,omitempty" xml:"fake_app_count,omitempty"`
	// 子任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 仿冒应用感染用户总数(任务完成时才返回)
	TotalInfectedUsers int64 `json:"total_infected_users,omitempty" xml:"total_infected_users,omitempty"`
}

var poolFakeAppSummary = sync.Pool{
	New: func() any {
		return new(FakeAppSummary)
	},
}

// GetFakeAppSummary() 从对象池中获取FakeAppSummary
func GetFakeAppSummary() *FakeAppSummary {
	return poolFakeAppSummary.Get().(*FakeAppSummary)
}

// ReleaseFakeAppSummary 释放FakeAppSummary
func ReleaseFakeAppSummary(v *FakeAppSummary) {
	v.FakeAppCount = 0
	v.Status = 0
	v.TotalInfectedUsers = 0
	poolFakeAppSummary.Put(v)
}
