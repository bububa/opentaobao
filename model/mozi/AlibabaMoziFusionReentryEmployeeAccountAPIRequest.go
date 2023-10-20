package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziFusionReentryEmployeeAccountAPIRequest 重新入职并且重新启用账号 API请求
// alibaba.mozi.fusion.reentry.employee.account
//
// 重新入职并且重新启用账号
type AlibabaMoziFusionReentryEmployeeAccountAPIRequest struct {
	model.Params
	// 入参
	_reentryEmployeeAccount *ReEntryTenantEmployeeAndAccountRequest
}

// NewAlibabaMoziFusionReentryEmployeeAccountRequest 初始化AlibabaMoziFusionReentryEmployeeAccountAPIRequest对象
func NewAlibabaMoziFusionReentryEmployeeAccountRequest() *AlibabaMoziFusionReentryEmployeeAccountAPIRequest {
	return &AlibabaMoziFusionReentryEmployeeAccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziFusionReentryEmployeeAccountAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.fusion.reentry.employee.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziFusionReentryEmployeeAccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziFusionReentryEmployeeAccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReentryEmployeeAccount is ReentryEmployeeAccount Setter
// 入参
func (r *AlibabaMoziFusionReentryEmployeeAccountAPIRequest) SetReentryEmployeeAccount(_reentryEmployeeAccount *ReEntryTenantEmployeeAndAccountRequest) error {
	r._reentryEmployeeAccount = _reentryEmployeeAccount
	r.Set("reentry.employee.account", _reentryEmployeeAccount)
	return nil
}

// GetReentryEmployeeAccount ReentryEmployeeAccount Getter
func (r AlibabaMoziFusionReentryEmployeeAccountAPIRequest) GetReentryEmployeeAccount() *ReEntryTenantEmployeeAndAccountRequest {
	return r._reentryEmployeeAccount
}
