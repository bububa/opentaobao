package charity

import (
	"sync"
)

// CharityTimeDto 结构体
type CharityTimeDto struct {
	// SPREAD_PUBLIC_WELFARE:公益传播,OFFLINE_PUBLIC_WELFARE:线下公益,PAN_PUBLIC_WELFARE:互联网公益,PUBLIC_DONATION::公益捐赠,WALK_DONATION:益起来,ENVIRONMENTAL_PROTECTION:ENVIRONMENTAL_PROTECTION
	CharityType string `json:"charity_type,omitempty" xml:"charity_type,omitempty"`
	// 公益类型名
	CharityTypeName string `json:"charity_type_name,omitempty" xml:"charity_type_name,omitempty"`
	// 获得公益时的时间
	ApproveDate string `json:"approve_date,omitempty" xml:"approve_date,omitempty"`
	// 标题
	MainTitle string `json:"main_title,omitempty" xml:"main_title,omitempty"`
	// 副标题
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 公益时产生时间
	GmtCreateTime string `json:"gmt_create_time,omitempty" xml:"gmt_create_time,omitempty"`
	// 记录id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 公益时，单位 分钟，6就是0.1公益时，60是1公益时
	CharityTime int64 `json:"charity_time,omitempty" xml:"charity_time,omitempty"`
	// 活动
	Activity *ActivityDto `json:"activity,omitempty" xml:"activity,omitempty"`
}

var poolCharityTimeDto = sync.Pool{
	New: func() any {
		return new(CharityTimeDto)
	},
}

// GetCharityTimeDto() 从对象池中获取CharityTimeDto
func GetCharityTimeDto() *CharityTimeDto {
	return poolCharityTimeDto.Get().(*CharityTimeDto)
}

// ReleaseCharityTimeDto 释放CharityTimeDto
func ReleaseCharityTimeDto(v *CharityTimeDto) {
	v.CharityType = ""
	v.CharityTypeName = ""
	v.ApproveDate = ""
	v.MainTitle = ""
	v.SubTitle = ""
	v.GmtCreateTime = ""
	v.Id = 0
	v.CharityTime = 0
	v.Activity = nil
	poolCharityTimeDto.Put(v)
}
