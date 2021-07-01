package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServiceSettleadjustmentCancelAPIRequest
取消结算调整单 API请求
tmall.service.settleadjustment.cancel

提供给服务商在对取消已经发起的结算调整单。
通过说明调整单ID进行结算调整单取消。 */
type TmallServiceSettleadjustmentCancelAPIRequest struct {
	model.Params
	// 结算调整单ID
	_id int64
	// 取消原因说明
	_comments string
}

// New
