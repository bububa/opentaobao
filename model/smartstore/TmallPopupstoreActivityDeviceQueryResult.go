package smartstore

import (
	"sync"
)

// TmallPopupstoreActivityDeviceQueryResult 结构体
type TmallPopupstoreActivityDeviceQueryResult struct {
	// 门店列表
	StoreList []Storelist `json:"store_list,omitempty" xml:"store_list>storelist,omitempty"`
	// 参与活动的品牌名，逗号分割
	SellerNames string `json:"seller_names,omitempty" xml:"seller_names,omitempty"`
	// 活动结束时间
	ActivityEndTime string `json:"activity_end_time,omitempty" xml:"activity_end_time,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 活动开始时间
	ActivityStartTime string `json:"activity_start_time,omitempty" xml:"activity_start_time,omitempty"`
	// 活动状态
	ActivityStatus int64 `json:"activity_status,omitempty" xml:"activity_status,omitempty"`
}

var poolTmallPopupstoreActivityDeviceQueryResult = sync.Pool{
	New: func() any {
		return new(TmallPopupstoreActivityDeviceQueryResult)
	},
}

// GetTmallPopupstoreActivityDeviceQueryResult() 从对象池中获取TmallPopupstoreActivityDeviceQueryResult
func GetTmallPopupstoreActivityDeviceQueryResult() *TmallPopupstoreActivityDeviceQueryResult {
	return poolTmallPopupstoreActivityDeviceQueryResult.Get().(*TmallPopupstoreActivityDeviceQueryResult)
}

// ReleaseTmallPopupstoreActivityDeviceQueryResult 释放TmallPopupstoreActivityDeviceQueryResult
func ReleaseTmallPopupstoreActivityDeviceQueryResult(v *TmallPopupstoreActivityDeviceQueryResult) {
	v.StoreList = v.StoreList[:0]
	v.SellerNames = ""
	v.ActivityEndTime = ""
	v.ActivityName = ""
	v.ActivityStartTime = ""
	v.ActivityStatus = 0
	poolTmallPopupstoreActivityDeviceQueryResult.Put(v)
}
