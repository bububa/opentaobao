package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkPosTradeReverseAPIRequest
轻pos品牌营销退款接口 API请求
alibaba.wdk.pos.trade.reverse

轻pos品牌营销场景，商家调用退款接口 */
type AlibabaWdkPosTradeReverseAPIRequest struct {
	model.Params
	// 退款请求
	_reverseRequest *FastBuyPosReverseRequest
}

// New
