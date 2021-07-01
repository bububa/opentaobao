package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressTradeBuyPlaceorderAPIRequest
AE下单API API请求
aliexpress.trade.buy.placeorder

A006_INVALID_ACCOUNT_INFO */
type AliexpressTradeBuyPlaceorderAPIRequest struct {
	model.Params
	// 下单具体参数
	_paramPlaceOrderRequest4OpenApiDTO *PlaceOrderRequest4OpenApiDto
}

// New
