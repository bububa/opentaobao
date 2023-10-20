package wdk

import (
	"sync"
)

// CommonActivityParam 结构体
type CommonActivityParam struct {
	// 自定义同步的渠道配置
	ChannelConfigList []ChannelConfig `json:"channel_config_list,omitempty" xml:"channel_config_list>channel_config,omitempty"`
	// 外部活动编码
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 活动Id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 是否自定义渠道同步
	ByChannel bool `json:"by_channel,omitempty" xml:"by_channel,omitempty"`
}

var poolCommonActivityParam = sync.Pool{
	New: func() any {
		return new(CommonActivityParam)
	},
}

// GetCommonActivityParam() 从对象池中获取CommonActivityParam
func GetCommonActivityParam() *CommonActivityParam {
	return poolCommonActivityParam.Get().(*CommonActivityParam)
}

// ReleaseCommonActivityParam 释放CommonActivityParam
func ReleaseCommonActivityParam(v *CommonActivityParam) {
	v.ChannelConfigList = v.ChannelConfigList[:0]
	v.OutActId = ""
	v.ActivityId = 0
	v.ByChannel = false
	poolCommonActivityParam.Put(v)
}
