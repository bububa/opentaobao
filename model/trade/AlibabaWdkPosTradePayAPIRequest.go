package trade

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkPosTradePayAPIRequest) Reset() {
	r._payRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkPosTradePayAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.pos.trade.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkPosTradePayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkPosTradePayAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaWdkPosTradePayAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkPosTradePayRequest()
	},
}

// GetAlibabaWdkPosTradePayRequest 从 sync.Pool 获取 AlibabaWdkPosTradePayAPIRequest
func GetAlibabaWdkPosTradePayAPIRequest() *AlibabaWdkPosTradePayAPIRequest {
	return poolAlibabaWdkPosTradePayAPIRequest.Get().(*AlibabaWdkPosTradePayAPIRequest)
}

// ReleaseAlibabaWdkPosTradePayAPIRequest 将 AlibabaWdkPosTradePayAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkPosTradePayAPIRequest(v *AlibabaWdkPosTradePayAPIRequest) {
	v.Reset()
	poolAlibabaWdkPosTradePayAPIRequest.Put(v)
}
