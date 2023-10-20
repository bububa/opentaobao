package wdk

import (
	"sync"
)

// CommonActivityRequest 结构体
type CommonActivityRequest struct {
	// 自定义同步的渠道配置
	ChannelConfigList []ChannelConfig `json:"channel_config_list,omitempty" xml:"channel_config_list>channel_config,omitempty"`
	// 商家活动id
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 五道口活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 是否自定义渠道同步
	ByChannel bool `json:"by_channel,omitempty" xml:"by_channel,omitempty"`
}

var poolCommonActivityRequest = sync.Pool{
	New: func() any {
		return new(CommonActivityRequest)
	},
}

// GetCommonActivityRequest() 从对象池中获取CommonActivityRequest
func GetCommonActivityRequest() *CommonActivityRequest {
	return poolCommonActivityRequest.Get().(*CommonActivityRequest)
}

// ReleaseCommonActivityRequest 释放CommonActivityRequest
func ReleaseCommonActivityRequest(v *CommonActivityRequest) {
	v.ChannelConfigList = v.ChannelConfigList[:0]
	v.OutActId = ""
	v.ActivityId = 0
	v.ByChannel = false
	poolCommonActivityRequest.Put(v)
}
