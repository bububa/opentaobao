package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkAxStoreCreateAPIRequest 翱象经营店创建接口 API请求
// alibaba.wdk.ax.store.create
//
// 翱象经营店创建
type AlibabaWdkAxStoreCreateAPIRequest struct {
	model.Params
	// 入参
	_axStoreCreateRequest *AxStoreCreateRequest
}

// NewAlibabaWdkAxStoreCreateRequest 初始化AlibabaWdkAxStoreCreateAPIRequest对象
func NewAlibabaWdkAxStoreCreateRequest() *AlibabaWdkAxStoreCreateAPIRequest {
	return &AlibabaWdkAxStoreCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkAxStoreCreateAPIRequest) Reset() {
	r._axStoreCreateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkAxStoreCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.ax.store.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkAxStoreCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkAxStoreCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAxStoreCreateRequest is AxStoreCreateRequest Setter
// 入参
func (r *AlibabaWdkAxStoreCreateAPIRequest) SetAxStoreCreateRequest(_axStoreCreateRequest *AxStoreCreateRequest) error {
	r._axStoreCreateRequest = _axStoreCreateRequest
	r.Set("ax_store_create_request", _axStoreCreateRequest)
	return nil
}

// GetAxStoreCreateRequest AxStoreCreateRequest Getter
func (r AlibabaWdkAxStoreCreateAPIRequest) GetAxStoreCreateRequest() *AxStoreCreateRequest {
	return r._axStoreCreateRequest
}

var poolAlibabaWdkAxStoreCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkAxStoreCreateRequest()
	},
}

// GetAlibabaWdkAxStoreCreateRequest 从 sync.Pool 获取 AlibabaWdkAxStoreCreateAPIRequest
func GetAlibabaWdkAxStoreCreateAPIRequest() *AlibabaWdkAxStoreCreateAPIRequest {
	return poolAlibabaWdkAxStoreCreateAPIRequest.Get().(*AlibabaWdkAxStoreCreateAPIRequest)
}

// ReleaseAlibabaWdkAxStoreCreateAPIRequest 将 AlibabaWdkAxStoreCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkAxStoreCreateAPIRequest(v *AlibabaWdkAxStoreCreateAPIRequest) {
	v.Reset()
	poolAlibabaWdkAxStoreCreateAPIRequest.Put(v)
}
