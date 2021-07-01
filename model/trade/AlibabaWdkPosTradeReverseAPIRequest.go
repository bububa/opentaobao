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

// NewAlibabaWdkPosTradeReverseRequest 初始化AlibabaWdkPosTradeReverseAPIRequest对象
func NewAlibabaWdkPosTradeReverseRequest() *AlibabaWdkPosTradeReverseAPIRequest {
	return &AlibabaWdkPosTradeReverseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkPosTradeReverseAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.pos.trade.reverse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkPosTradeReverseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ReverseRequest Setter
// 退款请求
func (r *AlibabaWdkPosTradeReverseAPIRequest) SetReverseRequest(_reverseRequest *FastBuyPosReverseRequest) error {
	r._reverseRequest = _reverseRequest
	r.Set("reverse_request", _reverseRequest)
	return nil
}

// Get ReverseRequest Getter
func (r AlibabaWdkPosTradeReverseAPIRequest) GetReverseRequest() *FastBuyPosReverseRequest {
	return r._reverseRequest
}
