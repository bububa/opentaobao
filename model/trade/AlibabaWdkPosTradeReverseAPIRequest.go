package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkpostradereverseAPIRequest 轻pos品牌营销退款接口 API请求
// alibaba.wdk.pos.trade.reverse
//
// 轻pos品牌营销场景，商家调用退款接口
type AlibabawdkpostradereverseAPIRequest struct {
	model.Params
	// 退款请求
	_reverseRequest *FastBuyPosReverseRequest
}

// NewAlibabawdkpostradereverseRequest 初始化AlibabawdkpostradereverseAPIRequest对象
func NewAlibabawdkpostradereverseRequest() *AlibabawdkpostradereverseAPIRequest {
	return &AlibabawdkpostradereverseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkpostradereverseAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.pos.trade.reverse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkpostradereverseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkpostradereverseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReverseRequest is ReverseRequest Setter
// 退款请求
func (r *AlibabawdkpostradereverseAPIRequest) SetReverseRequest(_reverseRequest *FastBuyPosReverseRequest) error {
	r._reverseRequest = _reverseRequest
	r.Set("reverse_request", _reverseRequest)
	return nil
}

// GetReverseRequest ReverseRequest Getter
func (r AlibabawdkpostradereverseAPIRequest) GetReverseRequest() *FastBuyPosReverseRequest {
	return r._reverseRequest
}
