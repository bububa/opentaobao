package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkPosTradeReverseAPIRequest 轻pos品牌营销退款接口 API请求
// alibaba.wdk.pos.trade.reverse
//
// 轻pos品牌营销场景，商家调用退款接口
type AlibabaWdkPosTradeReverseAPIRequest struct {
	model.Params
	// 退款请求
	_reverseRequest *FastBuyPosReverseRequest
}

// NewAlibabaWdkPosTradeReverseRequest 初始化AlibabaWdkPosTradeReverseAPIRequest对象
func NewAlibabaWdkPosTradeReverseRequest() *AlibabaWdkPosTradeReverseAPIRequest {
	return &AlibabaWdkPosTradeReverseAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkPosTradeReverseAPIRequest) Reset() {
	r._reverseRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkPosTradeReverseAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.pos.trade.reverse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkPosTradeReverseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkPosTradeReverseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReverseRequest is ReverseRequest Setter
// 退款请求
func (r *AlibabaWdkPosTradeReverseAPIRequest) SetReverseRequest(_reverseRequest *FastBuyPosReverseRequest) error {
	r._reverseRequest = _reverseRequest
	r.Set("reverse_request", _reverseRequest)
	return nil
}

// GetReverseRequest ReverseRequest Getter
func (r AlibabaWdkPosTradeReverseAPIRequest) GetReverseRequest() *FastBuyPosReverseRequest {
	return r._reverseRequest
}

var poolAlibabaWdkPosTradeReverseAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkPosTradeReverseRequest()
	},
}

// GetAlibabaWdkPosTradeReverseRequest 从 sync.Pool 获取 AlibabaWdkPosTradeReverseAPIRequest
func GetAlibabaWdkPosTradeReverseAPIRequest() *AlibabaWdkPosTradeReverseAPIRequest {
	return poolAlibabaWdkPosTradeReverseAPIRequest.Get().(*AlibabaWdkPosTradeReverseAPIRequest)
}

// ReleaseAlibabaWdkPosTradeReverseAPIRequest 将 AlibabaWdkPosTradeReverseAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkPosTradeReverseAPIRequest(v *AlibabaWdkPosTradeReverseAPIRequest) {
	v.Reset()
	poolAlibabaWdkPosTradeReverseAPIRequest.Put(v)
}
