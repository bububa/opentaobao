package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkPosTradeCloseAPIRequest
轻pos品牌营销关单接口 API请求
alibaba.wdk.pos.trade.close

轻pos品牌营销场景，提供关单接口给外部商家 */
type AlibabaWdkPosTradeCloseAPIRequest struct {
	model.Params
	// 关单请求
	_closeRequest *FastBuyPosCloseRequest
}

// New
