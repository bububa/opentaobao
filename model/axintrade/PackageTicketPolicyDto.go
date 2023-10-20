package axintrade

import (
	"sync"
)

// PackageTicketPolicyDto 结构体
type PackageTicketPolicyDto struct {
	// 使用规则名称
	TicketRuleName string `json:"ticket_rule_name,omitempty" xml:"ticket_rule_name,omitempty"`
	// 门票使用非结构化描述
	TicketDescription string `json:"ticket_description,omitempty" xml:"ticket_description,omitempty"`
	// 门票使用规则备注信息
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
}

var poolPackageTicketPolicyDto = sync.Pool{
	New: func() any {
		return new(PackageTicketPolicyDto)
	},
}

// GetPackageTicketPolicyDto() 从对象池中获取PackageTicketPolicyDto
func GetPackageTicketPolicyDto() *PackageTicketPolicyDto {
	return poolPackageTicketPolicyDto.Get().(*PackageTicketPolicyDto)
}

// ReleasePackageTicketPolicyDto 释放PackageTicketPolicyDto
func ReleasePackageTicketPolicyDto(v *PackageTicketPolicyDto) {
	v.TicketRuleName = ""
	v.TicketDescription = ""
	v.Remark = ""
	poolPackageTicketPolicyDto.Put(v)
}
