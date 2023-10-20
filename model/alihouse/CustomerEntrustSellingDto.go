package alihouse

import (
	"sync"
)

// CustomerEntrustSellingDto 结构体
type CustomerEntrustSellingDto struct {
	// 城市
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 委托小区名称
	EntrustCommunityName string `json:"entrust_community_name,omitempty" xml:"entrust_community_name,omitempty"`
	// 楼栋
	Building string `json:"building,omitempty" xml:"building,omitempty"`
	// 单元
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 房号
	HouseNumber string `json:"house_number,omitempty" xml:"house_number,omitempty"`
	// 称呼
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 联系电话
	PhoneNumber string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	// 户型
	HouseHold string `json:"house_hold,omitempty" xml:"house_hold,omitempty"`
	// 面积
	HouseSize string `json:"house_size,omitempty" xml:"house_size,omitempty"`
	// 装修
	Decoration string `json:"decoration,omitempty" xml:"decoration,omitempty"`
	// 联系信息
	ContactInfo string `json:"contact_info,omitempty" xml:"contact_info,omitempty"`
	// 业务模式（1-无忧卖房 2-普通卖房）
	BusiType int64 `json:"busi_type,omitempty" xml:"busi_type,omitempty"`
	// 小区ID
	OuterCommunityId int64 `json:"outer_community_id,omitempty" xml:"outer_community_id,omitempty"`
	// 期望售价单位 分
	ExpectingBidding int64 `json:"expecting_bidding,omitempty" xml:"expecting_bidding,omitempty"`
	// 室
	RoomNumber int64 `json:"room_number,omitempty" xml:"room_number,omitempty"`
	// 厅
	ParlorNumber int64 `json:"parlor_number,omitempty" xml:"parlor_number,omitempty"`
	// 卫
	ToiletNumber int64 `json:"toilet_number,omitempty" xml:"toilet_number,omitempty"`
	// 楼层 (-2 ----100)
	HouseFloor int64 `json:"house_floor,omitempty" xml:"house_floor,omitempty"`
	// 是否测试（0-否 1-是）
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}

var poolCustomerEntrustSellingDto = sync.Pool{
	New: func() any {
		return new(CustomerEntrustSellingDto)
	},
}

// GetCustomerEntrustSellingDto() 从对象池中获取CustomerEntrustSellingDto
func GetCustomerEntrustSellingDto() *CustomerEntrustSellingDto {
	return poolCustomerEntrustSellingDto.Get().(*CustomerEntrustSellingDto)
}

// ReleaseCustomerEntrustSellingDto 释放CustomerEntrustSellingDto
func ReleaseCustomerEntrustSellingDto(v *CustomerEntrustSellingDto) {
	v.CityCode = ""
	v.CityName = ""
	v.EntrustCommunityName = ""
	v.Building = ""
	v.Unit = ""
	v.HouseNumber = ""
	v.Name = ""
	v.PhoneNumber = ""
	v.HouseHold = ""
	v.HouseSize = ""
	v.Decoration = ""
	v.ContactInfo = ""
	v.BusiType = 0
	v.OuterCommunityId = 0
	v.ExpectingBidding = 0
	v.RoomNumber = 0
	v.ParlorNumber = 0
	v.ToiletNumber = 0
	v.HouseFloor = 0
	v.IsTest = 0
	poolCustomerEntrustSellingDto.Put(v)
}
