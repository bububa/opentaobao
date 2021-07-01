package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaBuynowOrderCreateAPIRequest
阿里巴巴买家buynow下单接口 API请求
alibaba.buynow.order.create

阿里巴巴买家下单接口 */
type AlibabaBuynowOrderCreateAPIRequest struct {
	model.Params
	// Order creation parameter
	_paramOrderCreateRequest *OrderCreateRequest
}

// New
