package guoguo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGuoguoCpBackupAssigncourierbyidAPIRequest
根据菜鸟账号ID指派小件员 API请求
cainiao.guoguo.cp.backup.assigncourierbyid

根据菜鸟账号ID指派小件员 */
type CainiaoGuoguoCpBackupAssigncourierbyidAPIRequest struct {
	model.Params
	// 指派/改派原因
	_assignReason string
	// 指派/改派原因编码
	_assignReasonCode string
	// 任务编号
	_taskId int64
	// 小件员菜鸟账号ID
	_accountId int64
	// CP公司编号
	_cpCode string
}

// New
