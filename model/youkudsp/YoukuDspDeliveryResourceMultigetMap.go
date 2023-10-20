package youkudsp

import (
	"sync"
)

// YoukuDspDeliveryResourceMultigetMap 结构体
type YoukuDspDeliveryResourceMultigetMap struct {
	// 图片json
	Imgs string `json:"imgs,omitempty" xml:"imgs,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 描述
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 换端地址
	DeeplinkUrl string `json:"deeplink_url,omitempty" xml:"deeplink_url,omitempty"`
	// 投放类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// h5地址
	DestinationUrl string `json:"destination_url,omitempty" xml:"destination_url,omitempty"`
}

var poolYoukuDspDeliveryResourceMultigetMap = sync.Pool{
	New: func() any {
		return new(YoukuDspDeliveryResourceMultigetMap)
	},
}

// GetYoukuDspDeliveryResourceMultigetMap() 从对象池中获取YoukuDspDeliveryResourceMultigetMap
func GetYoukuDspDeliveryResourceMultigetMap() *YoukuDspDeliveryResourceMultigetMap {
	return poolYoukuDspDeliveryResourceMultigetMap.Get().(*YoukuDspDeliveryResourceMultigetMap)
}

// ReleaseYoukuDspDeliveryResourceMultigetMap 释放YoukuDspDeliveryResourceMultigetMap
func ReleaseYoukuDspDeliveryResourceMultigetMap(v *YoukuDspDeliveryResourceMultigetMap) {
	v.Imgs = ""
	v.Title = ""
	v.Text = ""
	v.DeeplinkUrl = ""
	v.Type = ""
	v.DestinationUrl = ""
	poolYoukuDspDeliveryResourceMultigetMap.Put(v)
}
