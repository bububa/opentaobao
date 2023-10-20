package happytrip

import (
	"sync"
)

// TouristDto 结构体
type TouristDto struct {
	// 证件签发国
	CertNation string `json:"cert_nation,omitempty" xml:"cert_nation,omitempty"`
	// 姓
	FirstName string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 名
	LastName string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// 出行人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 主键
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 1男，2女
	Sex int64 `json:"sex,omitempty" xml:"sex,omitempty"`
	// 成人0，儿童1，婴儿2
	TouristType int64 `json:"tourist_type,omitempty" xml:"tourist_type,omitempty"`
	// 出行人0，同行人1，外部人员2
	TravelBusinessType int64 `json:"travel_business_type,omitempty" xml:"travel_business_type,omitempty"`
}

var poolTouristDto = sync.Pool{
	New: func() any {
		return new(TouristDto)
	},
}

// GetTouristDto() 从对象池中获取TouristDto
func GetTouristDto() *TouristDto {
	return poolTouristDto.Get().(*TouristDto)
}

// ReleaseTouristDto 释放TouristDto
func ReleaseTouristDto(v *TouristDto) {
	v.CertNation = ""
	v.FirstName = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.LastName = ""
	v.Name = ""
	v.UserId = ""
	v.Id = 0
	v.OrderId = 0
	v.Sex = 0
	v.TouristType = 0
	v.TravelBusinessType = 0
	poolTouristDto.Put(v)
}
