package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkPosTradePayAPIRequest
轻pos品牌营销支付接口 API请求
alibaba.wdk.pos.trade.pay

轻pos场景，外部商家支付后调用开放平台把支付信息回传给五道口交易 */
type AlibabaWdkPosTradePayAPIRequest struct {
	model.Params
	// 支付请求
	_payRequest *FastBuyPosPayRequest
}

// New
