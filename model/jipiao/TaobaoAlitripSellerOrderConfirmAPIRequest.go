package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripSellerOrderConfirmAPIRequest
代理商确认机票订单接口 API请求
taobao.alitrip.seller.order.confirm

此接口用于代理商确认机票订单。 */
type TaobaoAlitripSellerOrderConfirmAPIRequest struct {
	model.Params
	// 请求参数
	_tripConfirmOrderParam *TripConfirmOrderParam
}

// New
