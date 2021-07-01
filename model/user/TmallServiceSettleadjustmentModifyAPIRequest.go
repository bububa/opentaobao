package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServiceSettleadjustmentModifyAPIRequest
修改结算调整单 API请求
tmall.service.settleadjustment.modify

提供给服务商在对结算有异议时，发起结算调整单。
通过说明调整单ID，调整费用值，调整原因进行结算调整单修改。 */
type TmallServiceSettleadjustmentModifyAPIRequest struct {
	model.Params
	// 结算调整单父节点
	_paramSettleAdjustmentRequest *SettleAdjustmentRequest
}

// New
