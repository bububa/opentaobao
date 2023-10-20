package alihouse

import (
	"sync"
)

// SyncRichReviewDto 结构体
type SyncRichReviewDto struct {
	// 外部楼盘/小区id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部评测id
	OuterReviewId string `json:"outer_review_id,omitempty" xml:"outer_review_id,omitempty"`
	// 外部视频id
	OuterVideoId string `json:"outer_video_id,omitempty" xml:"outer_video_id,omitempty"`
	// 评测文本json数据
	ReviewJson string `json:"review_json,omitempty" xml:"review_json,omitempty"`
	// 发布时间
	PublishTime string `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	// 外部门店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 城市id
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 业务类型，1新房，2小区
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 评测状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 标识评测是否测试数据，0不是，1是
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}

var poolSyncRichReviewDto = sync.Pool{
	New: func() any {
		return new(SyncRichReviewDto)
	},
}

// GetSyncRichReviewDto() 从对象池中获取SyncRichReviewDto
func GetSyncRichReviewDto() *SyncRichReviewDto {
	return poolSyncRichReviewDto.Get().(*SyncRichReviewDto)
}

// ReleaseSyncRichReviewDto 释放SyncRichReviewDto
func ReleaseSyncRichReviewDto(v *SyncRichReviewDto) {
	v.OuterId = ""
	v.OuterReviewId = ""
	v.OuterVideoId = ""
	v.ReviewJson = ""
	v.PublishTime = ""
	v.OuterStoreId = ""
	v.CityId = 0
	v.BizType = 0
	v.Status = 0
	v.IsTest = 0
	poolSyncRichReviewDto.Put(v)
}
