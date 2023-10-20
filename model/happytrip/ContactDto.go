package happytrip

import (
	"sync"
)

// ContactDto 结构体
type ContactDto struct {
	// 联系人国家码
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 联系人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 是否默认联系人0不是，1是
	DefaultFlag int64 `json:"default_flag,omitempty" xml:"default_flag,omitempty"`
	// 主键
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 常用联系人ID
	RelationId int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 联系人的常用联系人id对应航旅常用联系人应用upc中的id,用于信息回流upc
	UpcId int64 `json:"upc_id,omitempty" xml:"upc_id,omitempty"`
}

var poolContactDto = sync.Pool{
	New: func() any {
		return new(ContactDto)
	},
}

// GetContactDto() 从对象池中获取ContactDto
func GetContactDto() *ContactDto {
	return poolContactDto.Get().(*ContactDto)
}

// ReleaseContactDto 释放ContactDto
func ReleaseContactDto(v *ContactDto) {
	v.CountryCode = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.Name = ""
	v.UserId = ""
	v.DefaultFlag = 0
	v.Id = 0
	v.OrderId = 0
	v.RelationId = 0
	v.UpcId = 0
	poolContactDto.Put(v)
}
