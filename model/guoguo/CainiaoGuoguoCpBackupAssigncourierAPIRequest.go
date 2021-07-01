package guoguo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGuoguoCpBackupAssigncourierAPIRequest
CP兜底后指定接单的小件员 API请求
cainiao.guoguo.cp.backup.assigncourier

CP兜底后指定接单的小件员；CP改派小件员 */
type CainiaoGuoguoCpBackupAssigncourierAPIRequest struct {
	model.Params
	// 小件员所在公司编号
	_cpCode string
	// 小件员员工编号
	_cpUserId string
	// LP订单号
	_lpCode string
	// 任务ID
	_taskId int64
	// 指派/改派原因编码
	_assignReasonCode string
	// 指派/改派原因
	_assignReason string
	// 小件员手机号
	_mobile string
}

// New
