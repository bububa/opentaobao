package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkPosTradeCreateAPIRequest
轻pos品牌营销下单接口 API请求
alibaba.wdk.pos.trade.create

提供给石基进行轻pos品牌营销下单 */
type AlibabaWdkPosTradeCreateAPIRequest struct {
	model.Params
	// 下单请求
	_createRequest *FastBuyPosCreateRequest
}

// New
