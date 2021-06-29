package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
人员离职 API请求
alibaba.mozi.fusion.dimission.employee.account

人员离职并且回收账号
*/
type AlibabaMoziFusionDimissionEmployeeAccountRequest struct {
    model.Params
    // 入参
    dimissionEmployee   *RemoveTenantEmployeeAndAccountRequest
}

// 初始化AlibabaMoziFusionDimissionEmployeeAccountRequest对象
func NewAlibabaMoziFusionDimissionEmployeeAccountRequest() *AlibabaMoziFusionDimissionEmployeeAccountRequest{
    return &AlibabaMoziFusionDimissionEmployeeAccountRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziFusionDimissionEmployeeAccountRequest) GetApiMethodName() string {
    return "alibaba.mozi.fusion.dimission.employee.account"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziFusionDimissionEmployeeAccountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DimissionEmployee Setter
// 入参
func (r *AlibabaMoziFusionDimissionEmployeeAccountRequest) SetDimissionEmployee(dimissionEmployee *RemoveTenantEmployeeAndAccountRequest) error {
    r.dimissionEmployee = dimissionEmployee
    r.Set("dimission.employee", dimissionEmployee)
    return nil
}

// DimissionEmployee Getter
func (r AlibabaMoziFusionDimissionEmployeeAccountRequest) GetDimissionEmployee() *RemoveTenantEmployeeAndAccountRequest {
    return r.dimissionEmployee
}
