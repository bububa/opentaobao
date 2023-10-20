package travel

import (
	"sync"
)

// ItineraryActivity 结构体
type ItineraryActivity struct {
	// 活动图片列表，多个图片以英文逗号分隔
	Images []string `json:"images,omitempty" xml:"images>string,omitempty"`
	// 活动标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 活动内容文本描述
	Txt string `json:"txt,omitempty" xml:"txt,omitempty"`
	// 活动预计时长，小时数
	Hour int64 `json:"hour,omitempty" xml:"hour,omitempty"`
	// 活动预计时长，分钟数
	Minute int64 `json:"minute,omitempty" xml:"minute,omitempty"`
}

var poolItineraryActivity = sync.Pool{
	New: func() any {
		return new(ItineraryActivity)
	},
}

// GetItineraryActivity() 从对象池中获取ItineraryActivity
func GetItineraryActivity() *ItineraryActivity {
	return poolItineraryActivity.Get().(*ItineraryActivity)
}

// ReleaseItineraryActivity 释放ItineraryActivity
func ReleaseItineraryActivity(v *ItineraryActivity) {
	v.Images = v.Images[:0]
	v.Title = ""
	v.Txt = ""
	v.Hour = 0
	v.Minute = 0
	poolItineraryActivity.Put(v)
}
