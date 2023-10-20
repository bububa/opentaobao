package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamozivdstenantapiservicetenantbyidAPIRequest 按租户ID查询租户信息 API请求
// alibaba.mozi.vds.tenant.api.service.tenantbyid
//
// 按租户ID查询租户信息
type AlibabamozivdstenantapiservicetenantbyidAPIRequest struct {
	model.Params
	// 入参
	_par0 *GetTenantByIdRequest
}

// NewAlibabamozivdstenantapiservicetenantbyidRequest 初始化AlibabamozivdstenantapiservicetenantbyidAPIRequest对象
func NewAlibabamozivdstenantapiservicetenantbyidRequest() *AlibabamozivdstenantapiservicetenantbyidAPIRequest {
	return &AlibabamozivdstenantapiservicetenantbyidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamozivdstenantapiservicetenantbyidAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.vds.tenant.api.service.tenantbyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamozivdstenantapiservicetenantbyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamozivdstenantapiservicetenantbyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPar0 is Par0 Setter
// 入参
func (r *AlibabamozivdstenantapiservicetenantbyidAPIRequest) SetPar0(_par0 *GetTenantByIdRequest) error {
	r._par0 = _par0
	r.Set("par0", _par0)
	return nil
}

// GetPar0 Par0 Getter
func (r AlibabamozivdstenantapiservicetenantbyidAPIRequest) GetPar0() *GetTenantByIdRequest {
	return r._par0
}
