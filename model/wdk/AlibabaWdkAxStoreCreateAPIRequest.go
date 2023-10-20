package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkaxstorecreateAPIRequest 翱象经营店创建接口 API请求
// alibaba.wdk.ax.store.create
//
// 翱象经营店创建
type AlibabawdkaxstorecreateAPIRequest struct {
	model.Params
	// 入参
	_axStoreCreateRequest *AxStoreCreateRequest
}

// NewAlibabawdkaxstorecreateRequest 初始化AlibabawdkaxstorecreateAPIRequest对象
func NewAlibabawdkaxstorecreateRequest() *AlibabawdkaxstorecreateAPIRequest {
	return &AlibabawdkaxstorecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkaxstorecreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ax.store.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkaxstorecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkaxstorecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAxStoreCreateRequest is AxStoreCreateRequest Setter
// 入参
func (r *AlibabawdkaxstorecreateAPIRequest) SetAxStoreCreateRequest(_axStoreCreateRequest *AxStoreCreateRequest) error {
	r._axStoreCreateRequest = _axStoreCreateRequest
	r.Set("ax_store_create_request", _axStoreCreateRequest)
	return nil
}

// GetAxStoreCreateRequest AxStoreCreateRequest Getter
func (r AlibabawdkaxstorecreateAPIRequest) GetAxStoreCreateRequest() *AxStoreCreateRequest {
	return r._axStoreCreateRequest
}
