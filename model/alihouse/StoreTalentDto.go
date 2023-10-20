package alihouse

import (
	"sync"
)

// StoreTalentDto 结构体
type StoreTalentDto struct {
	// 外部id
	TargetId string `json:"target_id,omitempty" xml:"target_id,omitempty"`
	// 达人号
	TalentId string `json:"talent_id,omitempty" xml:"talent_id,omitempty"`
	// 支付宝生活号id
	AlipayLifeIds string `json:"alipay_life_ids,omitempty" xml:"alipay_life_ids,omitempty"`
	// 账号类型
	AccountType int64 `json:"account_type,omitempty" xml:"account_type,omitempty"`
	// 是否换绑
	IsChangeBinding int64 `json:"is_change_binding,omitempty" xml:"is_change_binding,omitempty"`
	// 测试标
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
	// 1-绑定淘宝直播达人号  2-绑定支付宝直播主播号 默认为1（兼容，之后必填）
	BindType int64 `json:"bind_type,omitempty" xml:"bind_type,omitempty"`
}

var poolStoreTalentDto = sync.Pool{
	New: func() any {
		return new(StoreTalentDto)
	},
}

// GetStoreTalentDto() 从对象池中获取StoreTalentDto
func GetStoreTalentDto() *StoreTalentDto {
	return poolStoreTalentDto.Get().(*StoreTalentDto)
}

// ReleaseStoreTalentDto 释放StoreTalentDto
func ReleaseStoreTalentDto(v *StoreTalentDto) {
	v.TargetId = ""
	v.TalentId = ""
	v.AlipayLifeIds = ""
	v.AccountType = 0
	v.IsChangeBinding = 0
	v.IsTest = 0
	v.BindType = 0
	poolStoreTalentDto.Put(v)
}
