package smartstore

import (
	"sync"
)

// TmallPopupstoreActivityQueryResult 结构体
type TmallPopupstoreActivityQueryResult struct {
	// 活动结束时间
	ActivityEndTime string `json:"activity_end_time,omitempty" xml:"activity_end_time,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 活动开始时间
	ActivityStartTime string `json:"activity_start_time,omitempty" xml:"activity_start_time,omitempty"`
	// 活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 0：正常状态，-1：删除状态，-2：活动取消
	ActivityStatus int64 `json:"activity_status,omitempty" xml:"activity_status,omitempty"`
}

var poolTmallPopupstoreActivityQueryResult = sync.Pool{
	New: func() any {
		return new(TmallPopupstoreActivityQueryResult)
	},
}

// GetTmallPopupstoreActivityQueryResult() 从对象池中获取TmallPopupstoreActivityQueryResult
func GetTmallPopupstoreActivityQueryResult() *TmallPopupstoreActivityQueryResult {
	return poolTmallPopupstoreActivityQueryResult.Get().(*TmallPopupstoreActivityQueryResult)
}

// ReleaseTmallPopupstoreActivityQueryResult 释放TmallPopupstoreActivityQueryResult
func ReleaseTmallPopupstoreActivityQueryResult(v *TmallPopupstoreActivityQueryResult) {
	v.ActivityEndTime = ""
	v.ActivityName = ""
	v.ActivityStartTime = ""
	v.ActivityId = 0
	v.ActivityStatus = 0
	poolTmallPopupstoreActivityQueryResult.Put(v)
}
