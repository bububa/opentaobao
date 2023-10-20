package alihouse

import (
	"sync"
)

// ShopConfigDetailDto 结构体
type ShopConfigDetailDto struct {
	// 推荐经纪人排序
	ExcellentBroker []BrokerSortDto `json:"excellent_broker,omitempty" xml:"excellent_broker>broker_sort_dto,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 跳转地址
	JumpValue string `json:"jump_value,omitempty" xml:"jump_value,omitempty"`
	// 外部配置详细ID
	OuterConfigDetailId string `json:"outer_config_detail_id,omitempty" xml:"outer_config_detail_id,omitempty"`
	// 图片地址
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 标签名称
	Label string `json:"label,omitempty" xml:"label,omitempty"`
	// 副标题
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 跳转类型
	JumpType int64 `json:"jump_type,omitempty" xml:"jump_type,omitempty"`
	// 排序
	OrderNo int64 `json:"order_no,omitempty" xml:"order_no,omitempty"`
}

var poolShopConfigDetailDto = sync.Pool{
	New: func() any {
		return new(ShopConfigDetailDto)
	},
}

// GetShopConfigDetailDto() 从对象池中获取ShopConfigDetailDto
func GetShopConfigDetailDto() *ShopConfigDetailDto {
	return poolShopConfigDetailDto.Get().(*ShopConfigDetailDto)
}

// ReleaseShopConfigDetailDto 释放ShopConfigDetailDto
func ReleaseShopConfigDetailDto(v *ShopConfigDetailDto) {
	v.ExcellentBroker = v.ExcellentBroker[:0]
	v.Title = ""
	v.JumpValue = ""
	v.OuterConfigDetailId = ""
	v.ImageUrl = ""
	v.Label = ""
	v.SubTitle = ""
	v.JumpType = 0
	v.OrderNo = 0
	poolShopConfigDetailDto.Put(v)
}
