package aesolution

import (
	"sync"
)

// GlobalAeopTpIssueInfoDto 结构体
type GlobalAeopTpIssueInfoDto struct {
	// Issue model
	IssueModel string `json:"issue_model,omitempty" xml:"issue_model,omitempty"`
	// Issue status
	IssueStatus string `json:"issue_status,omitempty" xml:"issue_status,omitempty"`
	// issue creation time
	IssueTime string `json:"issue_time,omitempty" xml:"issue_time,omitempty"`
}

var poolGlobalAeopTpIssueInfoDto = sync.Pool{
	New: func() any {
		return new(GlobalAeopTpIssueInfoDto)
	},
}

// GetGlobalAeopTpIssueInfoDto() 从对象池中获取GlobalAeopTpIssueInfoDto
func GetGlobalAeopTpIssueInfoDto() *GlobalAeopTpIssueInfoDto {
	return poolGlobalAeopTpIssueInfoDto.Get().(*GlobalAeopTpIssueInfoDto)
}

// ReleaseGlobalAeopTpIssueInfoDto 释放GlobalAeopTpIssueInfoDto
func ReleaseGlobalAeopTpIssueInfoDto(v *GlobalAeopTpIssueInfoDto) {
	v.IssueModel = ""
	v.IssueStatus = ""
	v.IssueTime = ""
	poolGlobalAeopTpIssueInfoDto.Put(v)
}
