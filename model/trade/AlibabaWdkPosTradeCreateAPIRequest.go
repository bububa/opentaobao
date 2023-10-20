package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkpostradecreateAPIRequest 轻pos品牌营销下单接口 API请求
// alibaba.wdk.pos.trade.create
//
// 提供给石基进行轻pos品牌营销下单
type AlibabawdkpostradecreateAPIRequest struct {
	model.Params
	// 下单请求
	_createRequest *FastBuyPosCreateRequest
}

// NewAlibabawdkpostradecreateRequest 初始化AlibabawdkpostradecreateAPIRequest对象
func NewAlibabawdkpostradecreateRequest() *AlibabawdkpostradecreateAPIRequest {
	return &AlibabawdkpostradecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkpostradecreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.pos.trade.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkpostradecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkpostradecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateRequest is CreateRequest Setter
// 下单请求
func (r *AlibabawdkpostradecreateAPIRequest) SetCreateRequest(_createRequest *FastBuyPosCreateRequest) error {
	r._createRequest = _createRequest
	r.Set("create_request", _createRequest)
	return nil
}

// GetCreateRequest CreateRequest Getter
func (r AlibabawdkpostradecreateAPIRequest) GetCreateRequest() *FastBuyPosCreateRequest {
	return r._createRequest
}
