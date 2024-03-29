package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSocialLocaleGetAPIRequest Locale获取接口 API请求
// aliexpress.social.locale.get
//
// 新增Locale获取接口
type AliexpressSocialLocaleGetAPIRequest struct {
	model.Params
}

// NewAliexpressSocialLocaleGetRequest 初始化AliexpressSocialLocaleGetAPIRequest对象
func NewAliexpressSocialLocaleGetRequest() *AliexpressSocialLocaleGetAPIRequest {
	return &AliexpressSocialLocaleGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSocialLocaleGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSocialLocaleGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.social.locale.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSocialLocaleGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSocialLocaleGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAliexpressSocialLocaleGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSocialLocaleGetRequest()
	},
}

// GetAliexpressSocialLocaleGetRequest 从 sync.Pool 获取 AliexpressSocialLocaleGetAPIRequest
func GetAliexpressSocialLocaleGetAPIRequest() *AliexpressSocialLocaleGetAPIRequest {
	return poolAliexpressSocialLocaleGetAPIRequest.Get().(*AliexpressSocialLocaleGetAPIRequest)
}

// ReleaseAliexpressSocialLocaleGetAPIRequest 将 AliexpressSocialLocaleGetAPIRequest 放入 sync.Pool
func ReleaseAliexpressSocialLocaleGetAPIRequest(v *AliexpressSocialLocaleGetAPIRequest) {
	v.Reset()
	poolAliexpressSocialLocaleGetAPIRequest.Put(v)
}
