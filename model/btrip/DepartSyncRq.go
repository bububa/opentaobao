package btrip

import (
	"sync"
)

// DepartSyncRq 结构体
type DepartSyncRq struct {
	// 部门名称
	DepartName string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// 新版本第三方部门ID（和depart_id只传一个即可）
	ThirdDepartId string `json:"third_depart_id,omitempty" xml:"third_depart_id,omitempty"`
	// 新版本第三方部门父ID（和depart_pid只传一个即可）
	ThirdDepartPid string `json:"third_depart_pid,omitempty" xml:"third_depart_pid,omitempty"`
	// 部门主管id列表，多个主管id之前用&#39;|&#39;分割
	ManagerIds string `json:"manager_ids,omitempty" xml:"manager_ids,omitempty"`
	// 部门状态，0生效；1失效
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 上一版本第三方部门父部门ID（和third_depart_pid字段只传一个即可）
	DepartPid int64 `json:"depart_pid,omitempty" xml:"depart_pid,omitempty"`
	// 上一版本第三方部门ID（和third_depart_id只传一个即可）
	DepartId int64 `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
}

var poolDepartSyncRq = sync.Pool{
	New: func() any {
		return new(DepartSyncRq)
	},
}

// GetDepartSyncRq() 从对象池中获取DepartSyncRq
func GetDepartSyncRq() *DepartSyncRq {
	return poolDepartSyncRq.Get().(*DepartSyncRq)
}

// ReleaseDepartSyncRq 释放DepartSyncRq
func ReleaseDepartSyncRq(v *DepartSyncRq) {
	v.DepartName = ""
	v.ThirdDepartId = ""
	v.ThirdDepartPid = ""
	v.ManagerIds = ""
	v.Status = 0
	v.DepartPid = 0
	v.DepartId = 0
	poolDepartSyncRq.Put(v)
}
