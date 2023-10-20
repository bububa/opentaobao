package hotelhstdf

import (
	"sync"
)

// HotelRoomStaticDo 结构体
type HotelRoomStaticDo struct {
	// 字典code,收费停车场
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 数据银行id，酒店集团字典才有值
	DataBankBrandId string `json:"data_bank_brand_id,omitempty" xml:"data_bank_brand_id,omitempty"`
	// 描述，不一定有值
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 展示名称
	DisplayName string `json:"display_name,omitempty" xml:"display_name,omitempty"`
	// 英文名
	EnName string `json:"en_name,omitempty" xml:"en_name,omitempty"`
	// 暂不使用
	ExtendInfo string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
	// log图片
	LogoUrl string `json:"logo_url,omitempty" xml:"logo_url,omitempty"`
	// 中文名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 热度
	Hot int64 `json:"hot,omitempty" xml:"hot,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 父id
	ParentId int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
	// 优先级
	Priority int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// 状态，0--生效，-4--失效，接口返回一般均为生效数据
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 二级分类
	SubType int64 `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
	// 暂不使用
	Tag int64 `json:"tag,omitempty" xml:"tag,omitempty"`
	// 字典类型，10--预订须知，116--酒店设施，117--娱乐设施，118--酒店服务，119--房间设施
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolHotelRoomStaticDo = sync.Pool{
	New: func() any {
		return new(HotelRoomStaticDo)
	},
}

// GetHotelRoomStaticDo() 从对象池中获取HotelRoomStaticDo
func GetHotelRoomStaticDo() *HotelRoomStaticDo {
	return poolHotelRoomStaticDo.Get().(*HotelRoomStaticDo)
}

// ReleaseHotelRoomStaticDo 释放HotelRoomStaticDo
func ReleaseHotelRoomStaticDo(v *HotelRoomStaticDo) {
	v.Code = ""
	v.DataBankBrandId = ""
	v.Description = ""
	v.DisplayName = ""
	v.EnName = ""
	v.ExtendInfo = ""
	v.LogoUrl = ""
	v.Name = ""
	v.Hot = 0
	v.Id = 0
	v.ParentId = 0
	v.Priority = 0
	v.Status = 0
	v.SubType = 0
	v.Tag = 0
	v.Type = 0
	poolHotelRoomStaticDo.Put(v)
}
