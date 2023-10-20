package kclub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabakclubkcgetcategorytreeAPIRequest 知识云-查询租户下类目树 API请求
// alibaba.kclub.kc.getcategorytree
//
// 知识云-查询租户下类目树。通过租户id、类型(外部类目、帮助中心类目、内部类目)。
type AlibabakclubkcgetcategorytreeAPIRequest struct {
	model.Params
	// 租户id
	_tenantId int64
	// 鉴权参数
	_auth *TenancyAuth
}

// NewAlibabakclubkcgetcategorytreeRequest 初始化AlibabakclubkcgetcategorytreeAPIRequest对象
func NewAlibabakclubkcgetcategorytreeRequest() *AlibabakclubkcgetcategorytreeAPIRequest {
	return &AlibabakclubkcgetcategorytreeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabakclubkcgetcategorytreeAPIRequest) GetApiMethodName() string {
	return "alibaba.kclub.kc.getcategorytree"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabakclubkcgetcategorytreeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabakclubkcgetcategorytreeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantId is TenantId Setter
// 租户id
func (r *AlibabakclubkcgetcategorytreeAPIRequest) SetTenantId(_tenantId int64) error {
	r._tenantId = _tenantId
	r.Set("tenant_id", _tenantId)
	return nil
}

// GetTenantId TenantId Getter
func (r AlibabakclubkcgetcategorytreeAPIRequest) GetTenantId() int64 {
	return r._tenantId
}

// SetAuth is Auth Setter
// 鉴权参数
func (r *AlibabakclubkcgetcategorytreeAPIRequest) SetAuth(_auth *TenancyAuth) error {
	r._auth = _auth
	r.Set("auth", _auth)
	return nil
}

// GetAuth Auth Getter
func (r AlibabakclubkcgetcategorytreeAPIRequest) GetAuth() *TenancyAuth {
	return r._auth
}
