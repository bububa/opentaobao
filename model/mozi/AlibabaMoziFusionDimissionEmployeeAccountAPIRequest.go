package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziFusionDimissionEmployeeAccountAPIRequest 人员离职 API请求
// alibaba.mozi.fusion.dimission.employee.account
//
// 人员离职并且回收账号
type AlibabaMoziFusionDimissionEmployeeAccountAPIRequest struct {
	model.Params
	// 入参
	_dimissionEmployee *RemoveTenantEmployeeAndAccountRequest
}

// NewAlibabaMoziFusionDimissionEmployeeAccountRequest 初始化AlibabaMoziFusionDimissionEmployeeAccountAPIRequest对象
func NewAlibabaMoziFusionDimissionEmployeeAccountRequest() *AlibabaMoziFusionDimissionEmployeeAccountAPIRequest {
	return &AlibabaMoziFusionDimissionEmployeeAccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziFusionDimissionEmployeeAccountAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.fusion.dimission.employee.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziFusionDimissionEmployeeAccountAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DimissionEmployee Setter
// 入参
func (r *AlibabaMoziFusionDimissionEmployeeAccountAPIRequest) SetDimissionEmployee(_dimissionEmployee *RemoveTenantEmployeeAndAccountRequest) error {
	r._dimissionEmployee = _dimissionEmployee
	r.Set("dimission.employee", _dimissionEmployee)
	return nil
}

// Get DimissionEmployee Getter
func (r AlibabaMoziFusionDimissionEmployeeAccountAPIRequest) GetDimissionEmployee() *RemoveTenantEmployeeAndAccountRequest {
	return r._dimissionEmployee
}
