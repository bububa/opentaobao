package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdcAligeniusRefundsCheckAPIRequest
退款信息审核 API请求
taobao.rdc.aligenius.refunds.check

根据退款信息，对退款单进行审核 */
type TaobaoRdcAligeniusRefundsCheckAPIRequest struct {
	model.Params
	// 入参
	_param *RefundCheckDto
}

// New
