package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaServicecenterWorkcardCancelAPIRequest
服务平台工单取消接口 API请求
alibaba.servicecenter.workcard.cancel

取消服务工单 */
type AlibabaServicecenterWorkcardCancelAPIRequest struct {
	model.Params
	// 工单id
	_workcardId int64
	// 取消备注
	_memo string
	// 服务单id
	_serviceOrderId int64
	// 真实服务商昵称
	_realTpNick string
}

// New
