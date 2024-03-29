package drugtrace

import (
	"sync"
)

// CodeStatusTypeDto 结构体
type CodeStatusTypeDto struct {
	// 当前码的所有父码和packageLevel
	ParentCodeInfoList []ParentCodeInfo `json:"parent_code_info_list,omitempty" xml:"parent_code_info_list>parent_code_info,omitempty"`
	// 当前码的子码
	ChildCodes []Childcodes `json:"child_codes,omitempty" xml:"child_codes>childcodes,omitempty"`
	// 追溯码状态
	VerificationType string `json:"verification_type,omitempty" xml:"verification_type,omitempty"`
	// 最后业务日期
	LastBizDate string `json:"last_biz_date,omitempty" xml:"last_biz_date,omitempty"`
	// 码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 码状态
	CodeStatusNum string `json:"code_status_num,omitempty" xml:"code_status_num,omitempty"`
	// 码状态（A:已激活;I:已核注;O:已核注;C:已注销;E:码不存在）
	CodeStatus string `json:"code_status,omitempty" xml:"code_status,omitempty"`
	// 码状态描述
	CodeStatusDesc string `json:"code_status_desc,omitempty" xml:"code_status_desc,omitempty"`
	// 当前码的父码
	ParentCode string `json:"parent_code,omitempty" xml:"parent_code,omitempty"`
	// 最小包装数量
	CodeCount string `json:"code_count,omitempty" xml:"code_count,omitempty"`
	// 追溯码是否在本企业
	CurrEntStr string `json:"curr_ent_str,omitempty" xml:"curr_ent_str,omitempty"`
	// 追溯码状态描述
	VerificationTypeStr string `json:"verification_type_str,omitempty" xml:"verification_type_str,omitempty"`
	// 医保核销次数
	CheckCount int64 `json:"check_count,omitempty" xml:"check_count,omitempty"`
	// 父码下子码的数量
	IncludeProduceCodeCount int64 `json:"include_produce_code_count,omitempty" xml:"include_produce_code_count,omitempty"`
}

var poolCodeStatusTypeDto = sync.Pool{
	New: func() any {
		return new(CodeStatusTypeDto)
	},
}

// GetCodeStatusTypeDto() 从对象池中获取CodeStatusTypeDto
func GetCodeStatusTypeDto() *CodeStatusTypeDto {
	return poolCodeStatusTypeDto.Get().(*CodeStatusTypeDto)
}

// ReleaseCodeStatusTypeDto 释放CodeStatusTypeDto
func ReleaseCodeStatusTypeDto(v *CodeStatusTypeDto) {
	v.ParentCodeInfoList = v.ParentCodeInfoList[:0]
	v.ChildCodes = v.ChildCodes[:0]
	v.VerificationType = ""
	v.LastBizDate = ""
	v.Code = ""
	v.CodeStatusNum = ""
	v.CodeStatus = ""
	v.CodeStatusDesc = ""
	v.ParentCode = ""
	v.CodeCount = ""
	v.CurrEntStr = ""
	v.VerificationTypeStr = ""
	v.CheckCount = 0
	v.IncludeProduceCodeCount = 0
	poolCodeStatusTypeDto.Put(v)
}
