package alitripreceipt

import (
	"sync"
)

// IssueResultNotifyCmd 结构体
type IssueResultNotifyCmd struct {
	// 开票状态success:开票成功；fail:开票失败;issuing:开票中；cancel:已撤销；red：已冲红）
	IssueStatus string `json:"issue_status,omitempty" xml:"issue_status,omitempty"`
	// 开票流水ID
	IssueApplyId string `json:"issue_apply_id,omitempty" xml:"issue_apply_id,omitempty"`
	// 交易订单ID
	TpOrderId string `json:"tp_order_id,omitempty" xml:"tp_order_id,omitempty"`
	// 错误码
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 错误原因
	FailReason string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	// 是否最终
	IsFinally bool `json:"is_finally,omitempty" xml:"is_finally,omitempty"`
}

var poolIssueResultNotifyCmd = sync.Pool{
	New: func() any {
		return new(IssueResultNotifyCmd)
	},
}

// GetIssueResultNotifyCmd() 从对象池中获取IssueResultNotifyCmd
func GetIssueResultNotifyCmd() *IssueResultNotifyCmd {
	return poolIssueResultNotifyCmd.Get().(*IssueResultNotifyCmd)
}

// ReleaseIssueResultNotifyCmd 释放IssueResultNotifyCmd
func ReleaseIssueResultNotifyCmd(v *IssueResultNotifyCmd) {
	v.IssueStatus = ""
	v.IssueApplyId = ""
	v.TpOrderId = ""
	v.FailCode = ""
	v.FailReason = ""
	v.IsFinally = false
	poolIssueResultNotifyCmd.Put(v)
}
