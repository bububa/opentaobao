package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkAxStoreUpdateAPIRequest 翱翔经营店更新接口 API请求
// alibaba.wdk.ax.store.update
//
// 翱翔经营店更新接口
type AlibabaWdkAxStoreUpdateAPIRequest struct {
	model.Params
	// 入参
	_axStoreCreateRequest *AxStoreCreateRequest
}

// NewAlibabaWdkAxStoreUpdateRequest 初始化AlibabaWdkAxStoreUpdateAPIRequest对象
func NewAlibabaWdkAxStoreUpdateRequest() *AlibabaWdkAxStoreUpdateAPIRequest {
	return &AlibabaWdkAxStoreUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkAxStoreUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ax.store.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkAxStoreUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAxStoreCreateRequest is AxStoreCreateRequest Setter
// 入参
func (r *AlibabaWdkAxStoreUpdateAPIRequest) SetAxStoreCreateRequest(_axStoreCreateRequest *AxStoreCreateRequest) error {
	r._axStoreCreateRequest = _axStoreCreateRequest
	r.Set("ax_store_create_request", _axStoreCreateRequest)
	return nil
}

// GetAxStoreCreateRequest AxStoreCreateRequest Getter
func (r AlibabaWdkAxStoreUpdateAPIRequest) GetAxStoreCreateRequest() *AxStoreCreateRequest {
	return r._axStoreCreateRequest
}
