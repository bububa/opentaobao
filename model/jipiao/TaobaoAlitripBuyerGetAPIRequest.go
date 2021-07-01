package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripBuyerGetAPIRequest
敏感信息查询 API请求
taobao.alitrip.buyer.get

针对商家提供统一的TOP接口，可以根据订单获取订单对应买家联系电话（阿里小号）。 */
type TaobaoAlitripBuyerGetAPIRequest struct {
	model.Params
	// 敏感信息查询请求参数
	_requestAxb *RequestAxbDo
}

// New
