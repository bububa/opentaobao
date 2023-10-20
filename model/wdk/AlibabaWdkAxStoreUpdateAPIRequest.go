package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkAxStoreUpdateAPIRequest) Reset() {
	r._axStoreCreateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkAxStoreUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ax.store.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkAxStoreUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkAxStoreUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaWdkAxStoreUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkAxStoreUpdateRequest()
	},
}

// GetAlibabaWdkAxStoreUpdateRequest 从 sync.Pool 获取 AlibabaWdkAxStoreUpdateAPIRequest
func GetAlibabaWdkAxStoreUpdateAPIRequest() *AlibabaWdkAxStoreUpdateAPIRequest {
	return poolAlibabaWdkAxStoreUpdateAPIRequest.Get().(*AlibabaWdkAxStoreUpdateAPIRequest)
}

// ReleaseAlibabaWdkAxStoreUpdateAPIRequest 将 AlibabaWdkAxStoreUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkAxStoreUpdateAPIRequest(v *AlibabaWdkAxStoreUpdateAPIRequest) {
	v.Reset()
	poolAlibabaWdkAxStoreUpdateAPIRequest.Put(v)
}
