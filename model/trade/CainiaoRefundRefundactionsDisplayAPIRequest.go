package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoRefundRefundactionsDisplayAPIRequest
退货退款操作的展示信息(展现给买家) API请求
cainiao.refund.refundactions.display

退货退款操作的展示信息(展现给买家) */
type CainiaoRefundRefundactionsDisplayAPIRequest struct {
	model.Params
	// 请求入参
	_param0 *OrderRefundOperationReq
}

// New
