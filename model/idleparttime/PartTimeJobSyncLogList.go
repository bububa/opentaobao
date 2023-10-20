package idleparttime

import (
	"sync"
)

// PartTimeJobSyncLogList 结构体
type PartTimeJobSyncLogList struct {
	// 日志节点
	PartTimeJobTransNodes []PartTimeJobTransNodes `json:"part_time_job_trans_nodes,omitempty" xml:"part_time_job_trans_nodes>part_time_job_trans_nodes,omitempty"`
	// 当前状态描述
	CurrentStatus string `json:"current_status,omitempty" xml:"current_status,omitempty"`
	// 岗位id
	JobId int64 `json:"job_id,omitempty" xml:"job_id,omitempty"`
	// 岗位创建时间
	CreateTime int64 `json:"create_time,omitempty" xml:"create_time,omitempty"`
}

var poolPartTimeJobSyncLogList = sync.Pool{
	New: func() any {
		return new(PartTimeJobSyncLogList)
	},
}

// GetPartTimeJobSyncLogList() 从对象池中获取PartTimeJobSyncLogList
func GetPartTimeJobSyncLogList() *PartTimeJobSyncLogList {
	return poolPartTimeJobSyncLogList.Get().(*PartTimeJobSyncLogList)
}

// ReleasePartTimeJobSyncLogList 释放PartTimeJobSyncLogList
func ReleasePartTimeJobSyncLogList(v *PartTimeJobSyncLogList) {
	v.PartTimeJobTransNodes = v.PartTimeJobTransNodes[:0]
	v.CurrentStatus = ""
	v.JobId = 0
	v.CreateTime = 0
	poolPartTimeJobSyncLogList.Put(v)
}
