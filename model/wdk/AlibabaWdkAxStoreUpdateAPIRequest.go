package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkaxstoreupdateAPIRequest 翱翔经营店更新接口 API请求
// alibaba.wdk.ax.store.update
//
// 翱翔经营店更新接口
type AlibabawdkaxstoreupdateAPIRequest struct {
	model.Params
	// 入参
	_axStoreCreateRequest *AxStoreCreateRequest
}

// NewAlibabawdkaxstoreupdateRequest 初始化AlibabawdkaxstoreupdateAPIRequest对象
func NewAlibabawdkaxstoreupdateRequest() *AlibabawdkaxstoreupdateAPIRequest {
	return &AlibabawdkaxstoreupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkaxstoreupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ax.store.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkaxstoreupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkaxstoreupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAxStoreCreateRequest is AxStoreCreateRequest Setter
// 入参
func (r *AlibabawdkaxstoreupdateAPIRequest) SetAxStoreCreateRequest(_axStoreCreateRequest *AxStoreCreateRequest) error {
	r._axStoreCreateRequest = _axStoreCreateRequest
	r.Set("ax_store_create_request", _axStoreCreateRequest)
	return nil
}

// GetAxStoreCreateRequest AxStoreCreateRequest Getter
func (r AlibabawdkaxstoreupdateAPIRequest) GetAxStoreCreateRequest() *AxStoreCreateRequest {
	return r._axStoreCreateRequest
}
