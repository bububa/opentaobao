package kclub

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaKclubKcGetcategorytreeAPIRequest 知识云-查询租户下类目树 API请求
// alibaba.kclub.kc.getcategorytree
//
// 知识云-查询租户下类目树。通过租户id、类型(外部类目、帮助中心类目、内部类目)。
type AlibabaKclubKcGetcategorytreeAPIRequest struct {
	model.Params
	// 租户id
	_tenantId int64
	// 鉴权参数
	_auth *TenancyAuth
}

// NewAlibabaKclubKcGetcategorytreeRequest 初始化AlibabaKclubKcGetcategorytreeAPIRequest对象
func NewAlibabaKclubKcGetcategorytreeRequest() *AlibabaKclubKcGetcategorytreeAPIRequest {
	return &AlibabaKclubKcGetcategorytreeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaKclubKcGetcategorytreeAPIRequest) Reset() {
	r._tenantId = 0
	r._auth = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaKclubKcGetcategorytreeAPIRequest) GetApiMethodName() string {
	return "alibaba.kclub.kc.getcategorytree"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaKclubKcGetcategorytreeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaKclubKcGetcategorytreeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantId is TenantId Setter
// 租户id
func (r *AlibabaKclubKcGetcategorytreeAPIRequest) SetTenantId(_tenantId int64) error {
	r._tenantId = _tenantId
	r.Set("tenant_id", _tenantId)
	return nil
}

// GetTenantId TenantId Getter
func (r AlibabaKclubKcGetcategorytreeAPIRequest) GetTenantId() int64 {
	return r._tenantId
}

// SetAuth is Auth Setter
// 鉴权参数
func (r *AlibabaKclubKcGetcategorytreeAPIRequest) SetAuth(_auth *TenancyAuth) error {
	r._auth = _auth
	r.Set("auth", _auth)
	return nil
}

// GetAuth Auth Getter
func (r AlibabaKclubKcGetcategorytreeAPIRequest) GetAuth() *TenancyAuth {
	return r._auth
}

var poolAlibabaKclubKcGetcategorytreeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaKclubKcGetcategorytreeRequest()
	},
}

// GetAlibabaKclubKcGetcategorytreeRequest 从 sync.Pool 获取 AlibabaKclubKcGetcategorytreeAPIRequest
func GetAlibabaKclubKcGetcategorytreeAPIRequest() *AlibabaKclubKcGetcategorytreeAPIRequest {
	return poolAlibabaKclubKcGetcategorytreeAPIRequest.Get().(*AlibabaKclubKcGetcategorytreeAPIRequest)
}

// ReleaseAlibabaKclubKcGetcategorytreeAPIRequest 将 AlibabaKclubKcGetcategorytreeAPIRequest 放入 sync.Pool
func ReleaseAlibabaKclubKcGetcategorytreeAPIRequest(v *AlibabaKclubKcGetcategorytreeAPIRequest) {
	v.Reset()
	poolAlibabaKclubKcGetcategorytreeAPIRequest.Put(v)
}
