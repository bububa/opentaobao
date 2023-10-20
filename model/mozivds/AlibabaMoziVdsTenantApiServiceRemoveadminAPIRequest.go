package mozivds

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamozivdstenantapiserviceremoveadminAPIRequest 删除租户管理员服务 API请求
// alibaba.mozi.vds.tenant.api.service.removeadmin
//
// 删除租户管理员top服务
type AlibabamozivdstenantapiserviceremoveadminAPIRequest struct {
	model.Params
	// 请求入参
	_param *RemoveTenantAdminsRequest
}

// NewAlibabamozivdstenantapiserviceremoveadminRequest 初始化AlibabamozivdstenantapiserviceremoveadminAPIRequest对象
func NewAlibabamozivdstenantapiserviceremoveadminRequest() *AlibabamozivdstenantapiserviceremoveadminAPIRequest {
	return &AlibabamozivdstenantapiserviceremoveadminAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamozivdstenantapiserviceremoveadminAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.vds.tenant.api.service.removeadmin"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamozivdstenantapiserviceremoveadminAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamozivdstenantapiserviceremoveadminAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请求入参
func (r *AlibabamozivdstenantapiserviceremoveadminAPIRequest) SetParam(_param *RemoveTenantAdminsRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabamozivdstenantapiserviceremoveadminAPIRequest) GetParam() *RemoveTenantAdminsRequest {
	return r._param
}
