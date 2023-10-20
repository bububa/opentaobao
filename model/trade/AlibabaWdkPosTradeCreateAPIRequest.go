package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkPosTradeCreateAPIRequest 轻pos品牌营销下单接口 API请求
// alibaba.wdk.pos.trade.create
//
// 提供给石基进行轻pos品牌营销下单
type AlibabaWdkPosTradeCreateAPIRequest struct {
	model.Params
	// 下单请求
	_createRequest *FastBuyPosCreateRequest
}

// NewAlibabaWdkPosTradeCreateRequest 初始化AlibabaWdkPosTradeCreateAPIRequest对象
func NewAlibabaWdkPosTradeCreateRequest() *AlibabaWdkPosTradeCreateAPIRequest {
	return &AlibabaWdkPosTradeCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkPosTradeCreateAPIRequest) Reset() {
	r._createRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkPosTradeCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.pos.trade.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkPosTradeCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkPosTradeCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateRequest is CreateRequest Setter
// 下单请求
func (r *AlibabaWdkPosTradeCreateAPIRequest) SetCreateRequest(_createRequest *FastBuyPosCreateRequest) error {
	r._createRequest = _createRequest
	r.Set("create_request", _createRequest)
	return nil
}

// GetCreateRequest CreateRequest Getter
func (r AlibabaWdkPosTradeCreateAPIRequest) GetCreateRequest() *FastBuyPosCreateRequest {
	return r._createRequest
}

var poolAlibabaWdkPosTradeCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkPosTradeCreateRequest()
	},
}

// GetAlibabaWdkPosTradeCreateRequest 从 sync.Pool 获取 AlibabaWdkPosTradeCreateAPIRequest
func GetAlibabaWdkPosTradeCreateAPIRequest() *AlibabaWdkPosTradeCreateAPIRequest {
	return poolAlibabaWdkPosTradeCreateAPIRequest.Get().(*AlibabaWdkPosTradeCreateAPIRequest)
}

// ReleaseAlibabaWdkPosTradeCreateAPIRequest 将 AlibabaWdkPosTradeCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkPosTradeCreateAPIRequest(v *AlibabaWdkPosTradeCreateAPIRequest) {
	v.Reset()
	poolAlibabaWdkPosTradeCreateAPIRequest.Put(v)
}
