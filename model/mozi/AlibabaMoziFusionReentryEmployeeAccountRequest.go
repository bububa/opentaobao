package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
重新入职并且重新启用账号 API请求
alibaba.mozi.fusion.reentry.employee.account

重新入职并且重新启用账号
*/
type AlibabaMoziFusionReentryEmployeeAccountRequest struct {
    model.Params
    // 入参
    _reentryEmployeeAccount   *ReEntryTenantEmployeeAndAccountRequest
}

// 初始化AlibabaMoziFusionReentryEmployeeAccountRequest对象
func NewAlibabaMoziFusionReentryEmployeeAccountRequest() *AlibabaMoziFusionReentryEmployeeAccountRequest{
    return &AlibabaMoziFusionReentryEmployeeAccountRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziFusionReentryEmployeeAccountRequest) GetApiMethodName() string {
    return "alibaba.mozi.fusion.reentry.employee.account"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziFusionReentryEmployeeAccountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReentryEmployeeAccount Setter
// 入参
func (r *AlibabaMoziFusionReentryEmployeeAccountRequest) SetReentryEmployeeAccount(_reentryEmployeeAccount *ReEntryTenantEmployeeAndAccountRequest) error {
    r._reentryEmployeeAccount = _reentryEmployeeAccount
    r.Set("reentry.employee.account", _reentryEmployeeAccount)
    return nil
}

// ReentryEmployeeAccount Getter
func (r AlibabaMoziFusionReentryEmployeeAccountRequest) GetReentryEmployeeAccount() *ReEntryTenantEmployeeAndAccountRequest {
    return r._reentryEmployeeAccount
}
