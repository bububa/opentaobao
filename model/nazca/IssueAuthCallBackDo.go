package nazca

import (
	"sync"
)

// IssueAuthCallBackDo 结构体
type IssueAuthCallBackDo struct {
	// 合同编号
	ContractNum string `json:"contract_num,omitempty" xml:"contract_num,omitempty"`
	// 出证机构
	IssueOrg string `json:"issue_org,omitempty" xml:"issue_org,omitempty"`
	// 客户在1688的唯一标识
	PlatformUserId string `json:"platform_user_id,omitempty" xml:"platform_user_id,omitempty"`
	// 出证报告下载地址
	ReportUrl string `json:"report_url,omitempty" xml:"report_url,omitempty"`
	// 出证状态 0：出证中；1：出证成功
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

var poolIssueAuthCallBackDo = sync.Pool{
	New: func() any {
		return new(IssueAuthCallBackDo)
	},
}

// GetIssueAuthCallBackDo() 从对象池中获取IssueAuthCallBackDo
func GetIssueAuthCallBackDo() *IssueAuthCallBackDo {
	return poolIssueAuthCallBackDo.Get().(*IssueAuthCallBackDo)
}

// ReleaseIssueAuthCallBackDo 释放IssueAuthCallBackDo
func ReleaseIssueAuthCallBackDo(v *IssueAuthCallBackDo) {
	v.ContractNum = ""
	v.IssueOrg = ""
	v.PlatformUserId = ""
	v.ReportUrl = ""
	v.Status = ""
	poolIssueAuthCallBackDo.Put(v)
}
