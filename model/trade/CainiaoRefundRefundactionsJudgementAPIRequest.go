package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoRefundRefundactionsJudgementAPIRequest
判断当前用户是否能对订单执行一些逆向操作，比如退货操作 API请求
cainiao.refund.refundactions.judgement

判断当前用户是否能对订单执行一些逆向操作，比如退货操作 */
type CainiaoRefundRefundactionsJudgementAPIRequest struct {
	model.Params
	// 操作请求
	_param0 *OrderRefundOperationJudgementReq
}

// New
