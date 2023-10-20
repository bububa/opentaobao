package tbk

import (
	"sync"
)

// TaobaoTbkScPublisherInfoGetData 结构体
type TaobaoTbkScPublisherInfoGetData struct {
	// 共享字段 - 渠道或会员列表
	InviterList []TaobaoTbkScPublisherInfoGetMapData `json:"inviter_list,omitempty" xml:"inviter_list>taobao_tbk_sc_publisher_info_get_map_data,omitempty"`
	// 渠道专属pidList
	RootPidChannelList []string `json:"root_pid_channel_list,omitempty" xml:"root_pid_channel_list>string,omitempty"`
	// 共享字段 - 总记录数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolTaobaoTbkScPublisherInfoGetData = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScPublisherInfoGetData)
	},
}

// GetTaobaoTbkScPublisherInfoGetData() 从对象池中获取TaobaoTbkScPublisherInfoGetData
func GetTaobaoTbkScPublisherInfoGetData() *TaobaoTbkScPublisherInfoGetData {
	return poolTaobaoTbkScPublisherInfoGetData.Get().(*TaobaoTbkScPublisherInfoGetData)
}

// ReleaseTaobaoTbkScPublisherInfoGetData 释放TaobaoTbkScPublisherInfoGetData
func ReleaseTaobaoTbkScPublisherInfoGetData(v *TaobaoTbkScPublisherInfoGetData) {
	v.InviterList = v.InviterList[:0]
	v.RootPidChannelList = v.RootPidChannelList[:0]
	v.TotalCount = 0
	poolTaobaoTbkScPublisherInfoGetData.Put(v)
}
