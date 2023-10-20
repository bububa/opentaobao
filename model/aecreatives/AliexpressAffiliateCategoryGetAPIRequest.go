package aecreatives

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateCategoryGetAPIRequest AE流量推广类目信息获取API API请求
// aliexpress.affiliate.category.get
//
// 获取AE流量推广类目的API
type AliexpressAffiliateCategoryGetAPIRequest struct {
	model.Params
	// 请求安全签名
	_appSignature string
}

// NewAliexpressAffiliateCategoryGetRequest 初始化AliexpressAffiliateCategoryGetAPIRequest对象
func NewAliexpressAffiliateCategoryGetRequest() *AliexpressAffiliateCategoryGetAPIRequest {
	return &AliexpressAffiliateCategoryGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressAffiliateCategoryGetAPIRequest) Reset() {
	r._appSignature = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateCategoryGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.category.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateCategoryGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAffiliateCategoryGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppSignature is AppSignature Setter
// 请求安全签名
func (r *AliexpressAffiliateCategoryGetAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// GetAppSignature AppSignature Getter
func (r AliexpressAffiliateCategoryGetAPIRequest) GetAppSignature() string {
	return r._appSignature
}

var poolAliexpressAffiliateCategoryGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressAffiliateCategoryGetRequest()
	},
}

// GetAliexpressAffiliateCategoryGetRequest 从 sync.Pool 获取 AliexpressAffiliateCategoryGetAPIRequest
func GetAliexpressAffiliateCategoryGetAPIRequest() *AliexpressAffiliateCategoryGetAPIRequest {
	return poolAliexpressAffiliateCategoryGetAPIRequest.Get().(*AliexpressAffiliateCategoryGetAPIRequest)
}

// ReleaseAliexpressAffiliateCategoryGetAPIRequest 将 AliexpressAffiliateCategoryGetAPIRequest 放入 sync.Pool
func ReleaseAliexpressAffiliateCategoryGetAPIRequest(v *AliexpressAffiliateCategoryGetAPIRequest) {
	v.Reset()
	poolAliexpressAffiliateCategoryGetAPIRequest.Put(v)
}
