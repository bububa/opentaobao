package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServiceSettleadjustmentRequestAPIRequest
创建结算调整单 API请求
tmall.service.settleadjustment.request

提供给服务商在对结算有异议时，发起结算调整单。
通过说明工单ID，调整费用值，调整原因进行新建结算调整单。 */
type TmallServiceSettleadjustmentRequestAPIRequest struct {
	model.Params
	// 父节点
	_paramSettleAdjustmentRequest *SettleAdjustmentRequest
}

// New
