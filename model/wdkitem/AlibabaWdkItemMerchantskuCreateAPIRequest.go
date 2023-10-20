package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemmerchantskucreateAPIRequest 商家商品信息新建 API请求
// alibaba.wdk.item.merchantsku.create
//
// 商家商品信息新建
type AlibabawdkitemmerchantskucreateAPIRequest struct {
	model.Params
	// 新建商品参数，是个json字符串
	_params string
}

// NewAlibabawdkitemmerchantskucreateRequest 初始化AlibabawdkitemmerchantskucreateAPIRequest对象
func NewAlibabawdkitemmerchantskucreateRequest() *AlibabawdkitemmerchantskucreateAPIRequest {
	return &AlibabawdkitemmerchantskucreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkitemmerchantskucreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.merchantsku.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkitemmerchantskucreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkitemmerchantskucreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParams is Params Setter
// 新建商品参数，是个json字符串
func (r *AlibabawdkitemmerchantskucreateAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r AlibabawdkitemmerchantskucreateAPIRequest) GetParams() string {
	return r._params
}
