package tbtrade

import (
	"sync"
)

// MonitorUrlTopRequestDto 结构体
type MonitorUrlTopRequestDto struct {
	// 自研页的H5链接；若落地页类型为“自研页”，该项必填
	CusH5Url string `json:"cus_h5_url,omitempty" xml:"cus_h5_url,omitempty"`
	// 需要封装成唤端链接的原始https链接；若落地页类型为“自研页”，该项必填
	CusDeepLink string `json:"cus_deep_link,omitempty" xml:"cus_deep_link,omitempty"`
	// 监测唯一id，首次创建不需要传，通过出参返回。修改监测落地页时必填
	MonitorUniqueId string `json:"monitor_unique_id,omitempty" xml:"monitor_unique_id,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 落地页类型：1-商品详情页；2-自研页
	PageType int64 `json:"page_type,omitempty" xml:"page_type,omitempty"`
	// 渠道id：1-巨量2.0，104-腾讯3.0，105-快手，108-B站
	ChannelId int64 `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// 产品类型：1-udsmart 2-流量通Pro，不传默认为udsmart
	ProdType int64 `json:"prod_type,omitempty" xml:"prod_type,omitempty"`
}

var poolMonitorUrlTopRequestDto = sync.Pool{
	New: func() any {
		return new(MonitorUrlTopRequestDto)
	},
}

// GetMonitorUrlTopRequestDto() 从对象池中获取MonitorUrlTopRequestDto
func GetMonitorUrlTopRequestDto() *MonitorUrlTopRequestDto {
	return poolMonitorUrlTopRequestDto.Get().(*MonitorUrlTopRequestDto)
}

// ReleaseMonitorUrlTopRequestDto 释放MonitorUrlTopRequestDto
func ReleaseMonitorUrlTopRequestDto(v *MonitorUrlTopRequestDto) {
	v.CusH5Url = ""
	v.CusDeepLink = ""
	v.MonitorUniqueId = ""
	v.ItemId = 0
	v.PageType = 0
	v.ChannelId = 0
	v.ProdType = 0
	poolMonitorUrlTopRequestDto.Put(v)
}
