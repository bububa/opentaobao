package idleparttime

import (
	"sync"
)

// AlibabaIdleParttimeSynclogData 结构体
type AlibabaIdleParttimeSynclogData struct {
	// 岗位列表
	PartTimeJobSyncLogList []PartTimeJobSyncLogList `json:"part_time_job_sync_log_list,omitempty" xml:"part_time_job_sync_log_list>part_time_job_sync_log_list,omitempty"`
}

var poolAlibabaIdleParttimeSynclogData = sync.Pool{
	New: func() any {
		return new(AlibabaIdleParttimeSynclogData)
	},
}

// GetAlibabaIdleParttimeSynclogData() 从对象池中获取AlibabaIdleParttimeSynclogData
func GetAlibabaIdleParttimeSynclogData() *AlibabaIdleParttimeSynclogData {
	return poolAlibabaIdleParttimeSynclogData.Get().(*AlibabaIdleParttimeSynclogData)
}

// ReleaseAlibabaIdleParttimeSynclogData 释放AlibabaIdleParttimeSynclogData
func ReleaseAlibabaIdleParttimeSynclogData(v *AlibabaIdleParttimeSynclogData) {
	v.PartTimeJobSyncLogList = v.PartTimeJobSyncLogList[:0]
	poolAlibabaIdleParttimeSynclogData.Put(v)
}
