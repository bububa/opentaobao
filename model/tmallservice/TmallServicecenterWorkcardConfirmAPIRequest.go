package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardConfirmAPIRequest
服务商确认服务完成 API请求
tmall.servicecenter.workcard.confirm

提供给外部合作服务商，用于通知天猫，告知寄修服务厂内操作全部完成 */
type TmallServicecenterWorkcardConfirmAPIRequest struct {
	model.Params
	// 工单id
	_workcardId int64
}

// New
