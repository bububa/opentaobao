package wlbimports

import (
	"sync"
)

// IdentityItemDto 结构体
type IdentityItemDto struct {
	// 防伪扣编码
	UniCode string `json:"uni_code,omitempty" xml:"uni_code,omitempty"`
	// 货品Id
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 鉴定结果备注
	IdentityRemark string `json:"identity_remark,omitempty" xml:"identity_remark,omitempty"`
	// 鉴定报告地址
	ReportUrl string `json:"report_url,omitempty" xml:"report_url,omitempty"`
	// 鉴定结果
	IdentityResult string `json:"identity_result,omitempty" xml:"identity_result,omitempty"`
	// 鉴定次数
	IdentityCnt int64 `json:"identity_cnt,omitempty" xml:"identity_cnt,omitempty"`
}

var poolIdentityItemDto = sync.Pool{
	New: func() any {
		return new(IdentityItemDto)
	},
}

// GetIdentityItemDto() 从对象池中获取IdentityItemDto
func GetIdentityItemDto() *IdentityItemDto {
	return poolIdentityItemDto.Get().(*IdentityItemDto)
}

// ReleaseIdentityItemDto 释放IdentityItemDto
func ReleaseIdentityItemDto(v *IdentityItemDto) {
	v.UniCode = ""
	v.ScItemId = ""
	v.IdentityRemark = ""
	v.ReportUrl = ""
	v.IdentityResult = ""
	v.IdentityCnt = 0
	poolIdentityItemDto.Put(v)
}
