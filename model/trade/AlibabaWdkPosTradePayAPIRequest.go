package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkpostradepayAPIRequest 轻pos品牌营销支付接口 API请求
// alibaba.wdk.pos.trade.pay
//
// 轻pos场景，外部商家支付后调用开放平台把支付信息回传给五道口交易
type AlibabawdkpostradepayAPIRequest struct {
	model.Params
	// 支付请求
	_payRequest *FastBuyPosPayRequest
}

// NewAlibabawdkpostradepayRequest 初始化AlibabawdkpostradepayAPIRequest对象
func NewAlibabawdkpostradepayRequest() *AlibabawdkpostradepayAPIRequest {
	return &AlibabawdkpostradepayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkpostradepayAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.pos.trade.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkpostradepayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkpostradepayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPayRequest is PayRequest Setter
// 支付请求
func (r *AlibabawdkpostradepayAPIRequest) SetPayRequest(_payRequest *FastBuyPosPayRequest) error {
	r._payRequest = _payRequest
	r.Set("pay_request", _payRequest)
	return nil
}

// GetPayRequest PayRequest Getter
func (r AlibabawdkpostradepayAPIRequest) GetPayRequest() *FastBuyPosPayRequest {
	return r._payRequest
}
