package tmallnr

import (
	"sync"
)

// NrtCrmActivityStatisticsDataReq 结构体
type NrtCrmActivityStatisticsDataReq struct {
	// 数据上传时间
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 业务身份
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 浏览人数
	Uv int64 `json:"uv,omitempty" xml:"uv,omitempty"`
	// 活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 页面访问数
	Pv int64 `json:"pv,omitempty" xml:"pv,omitempty"`
	// 摊位id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 导购员id
	GuiderId int64 `json:"guider_id,omitempty" xml:"guider_id,omitempty"`
}

var poolNrtCrmActivityStatisticsDataReq = sync.Pool{
	New: func() any {
		return new(NrtCrmActivityStatisticsDataReq)
	},
}

// GetNrtCrmActivityStatisticsDataReq() 从对象池中获取NrtCrmActivityStatisticsDataReq
func GetNrtCrmActivityStatisticsDataReq() *NrtCrmActivityStatisticsDataReq {
	return poolNrtCrmActivityStatisticsDataReq.Get().(*NrtCrmActivityStatisticsDataReq)
}

// ReleaseNrtCrmActivityStatisticsDataReq 释放NrtCrmActivityStatisticsDataReq
func ReleaseNrtCrmActivityStatisticsDataReq(v *NrtCrmActivityStatisticsDataReq) {
	v.Date = ""
	v.BizCode = ""
	v.Uv = 0
	v.ActivityId = 0
	v.Pv = 0
	v.StoreId = 0
	v.GuiderId = 0
	poolNrtCrmActivityStatisticsDataReq.Put(v)
}
