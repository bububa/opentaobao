package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardVerifycodeResendAPIRequest
重发核销码 API请求
tmall.servicecenter.workcard.verifycode.resend

重发核销码 */
type TmallServicecenterWorkcardVerifycodeResendAPIRequest struct {
	model.Params
	// 工单id
	_workcardId int64
	// 门店/网点id
	_serviceStoreId int64
}

// New
