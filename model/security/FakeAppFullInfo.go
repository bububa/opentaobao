package security

import (
	"sync"
)

// FakeAppFullInfo 结构体
type FakeAppFullInfo struct {
	// 仿冒应用列表(任务完成时才返回)
	FakeAppDetails []FakeAppDetail `json:"fake_app_details,omitempty" xml:"fake_app_details>fake_app_detail,omitempty"`
	// 仿冒应用总数(任务完成时才返回)
	FakeAppCount int64 `json:"fake_app_count,omitempty" xml:"fake_app_count,omitempty"`
	// 子任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 仿冒应用感染用户总数(任务完成时才返回)
	TotalInfectedUsers int64 `json:"total_infected_users,omitempty" xml:"total_infected_users,omitempty"`
}

var poolFakeAppFullInfo = sync.Pool{
	New: func() any {
		return new(FakeAppFullInfo)
	},
}

// GetFakeAppFullInfo() 从对象池中获取FakeAppFullInfo
func GetFakeAppFullInfo() *FakeAppFullInfo {
	return poolFakeAppFullInfo.Get().(*FakeAppFullInfo)
}

// ReleaseFakeAppFullInfo 释放FakeAppFullInfo
func ReleaseFakeAppFullInfo(v *FakeAppFullInfo) {
	v.FakeAppDetails = v.FakeAppDetails[:0]
	v.FakeAppCount = 0
	v.Status = 0
	v.TotalInfectedUsers = 0
	poolFakeAppFullInfo.Put(v)
}
