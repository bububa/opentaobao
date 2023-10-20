package crm

import (
	"sync"
)

// TaobaoMemberHismemberdataGetPrivyResultDto 结构体
type TaobaoMemberHismemberdataGetPrivyResultDto struct {
	// 备份会员信息列表
	HsmemberinfoList []HsMemberInfoDto `json:"hsmemberinfo_list,omitempty" xml:"hsmemberinfo_list>hs_member_info_dto,omitempty"`
	// 总记录数，分页用
	Total string `json:"total,omitempty" xml:"total,omitempty"`
	// 返回信息，成功时可为空
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回码，0：表示查询成功，其他：表示失败
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}

var poolTaobaoMemberHismemberdataGetPrivyResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoMemberHismemberdataGetPrivyResultDto)
	},
}

// GetTaobaoMemberHismemberdataGetPrivyResultDto() 从对象池中获取TaobaoMemberHismemberdataGetPrivyResultDto
func GetTaobaoMemberHismemberdataGetPrivyResultDto() *TaobaoMemberHismemberdataGetPrivyResultDto {
	return poolTaobaoMemberHismemberdataGetPrivyResultDto.Get().(*TaobaoMemberHismemberdataGetPrivyResultDto)
}

// ReleaseTaobaoMemberHismemberdataGetPrivyResultDto 释放TaobaoMemberHismemberdataGetPrivyResultDto
func ReleaseTaobaoMemberHismemberdataGetPrivyResultDto(v *TaobaoMemberHismemberdataGetPrivyResultDto) {
	v.HsmemberinfoList = v.HsmemberinfoList[:0]
	v.Total = ""
	v.Msg = ""
	v.Code = ""
	poolTaobaoMemberHismemberdataGetPrivyResultDto.Put(v)
}
