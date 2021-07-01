package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripIeBuyerOrderBookpayAPIRequest
【国际机票】下单预定支付 API请求
alitrip.ie.buyer.order.bookpay

【国际机票】 生单预定支付接口 */
type AlitripIeBuyerOrderBookpayAPIRequest struct {
	model.Params
	// 生单支付请求参数
	_bookPayOrderParam *BookPayOrderRq
}

// New
