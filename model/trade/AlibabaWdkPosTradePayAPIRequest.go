package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkPosTradePayAPIRequest 轻pos品牌营销支付接口 API请求
// alibaba.wdk.pos.trade.pay
//
// 轻pos场景，外部商家支付后调用开放平台把支付信息回传给五道口交易
type AlibabaWdkPosTradePayAPIRequest struct {
	model.Params
	// 支付请求
	_payRequest *FastBuyPosPayRequest
}

// NewAlibabaWdkPosTradePayRequest 初始化AlibabaWdkPosTradePayAPIRequest对象
func NewAlibabaWdkPosTradePayRequest() *AlibabaWdkPosTradePayAPIRequest {
	return &AlibabaWdkPosTradePayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkPosTradePayAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.pos.trade.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkPosTradePayAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPayRequest is PayRequest Setter
// 支付请求
func (r *AlibabaWdkPosTradePayAPIRequest) SetPayRequest(_payRequest *FastBuyPosPayRequest) error {
	r._payRequest = _payRequest
	r.Set("pay_request", _payRequest)
	return nil
}

// GetPayRequest PayRequest Getter
func (r AlibabaWdkPosTradePayAPIRequest) GetPayRequest() *FastBuyPosPayRequest {
	return r._payRequest
}
