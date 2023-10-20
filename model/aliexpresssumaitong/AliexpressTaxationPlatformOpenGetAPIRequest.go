package aliexpresssumaitong

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressTaxationPlatformOpenGetAPIRequest 平台固定参数获取 API请求
// aliexpress.taxation.platform.open.get
//
// Aliexpress开放平台固定参数获取
type AliexpressTaxationPlatformOpenGetAPIRequest struct {
	model.Params
}

// NewAliexpressTaxationPlatformOpenGetRequest 初始化AliexpressTaxationPlatformOpenGetAPIRequest对象
func NewAliexpressTaxationPlatformOpenGetRequest() *AliexpressTaxationPlatformOpenGetAPIRequest {
	return &AliexpressTaxationPlatformOpenGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressTaxationPlatformOpenGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressTaxationPlatformOpenGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.taxation.platform.open.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressTaxationPlatformOpenGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressTaxationPlatformOpenGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAliexpressTaxationPlatformOpenGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressTaxationPlatformOpenGetRequest()
	},
}

// GetAliexpressTaxationPlatformOpenGetRequest 从 sync.Pool 获取 AliexpressTaxationPlatformOpenGetAPIRequest
func GetAliexpressTaxationPlatformOpenGetAPIRequest() *AliexpressTaxationPlatformOpenGetAPIRequest {
	return poolAliexpressTaxationPlatformOpenGetAPIRequest.Get().(*AliexpressTaxationPlatformOpenGetAPIRequest)
}

// ReleaseAliexpressTaxationPlatformOpenGetAPIRequest 将 AliexpressTaxationPlatformOpenGetAPIRequest 放入 sync.Pool
func ReleaseAliexpressTaxationPlatformOpenGetAPIRequest(v *AliexpressTaxationPlatformOpenGetAPIRequest) {
	v.Reset()
	poolAliexpressTaxationPlatformOpenGetAPIRequest.Put(v)
}
