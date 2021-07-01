package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeaseTailpaymentbackAPIRequest
尾款处置方案回传 API请求
tmall.car.lease.tailpaymentback

尾款处置方案回传 */
type TmallCarLeaseTailpaymentbackAPIRequest struct {
	model.Params
	// 尾款方案
	_tailPaymentDTO *TailPaymentDto
}

// New
