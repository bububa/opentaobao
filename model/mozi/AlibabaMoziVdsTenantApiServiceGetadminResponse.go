package mozi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取员工租户管理员信息（查询员工是否为租户管理员） API返回值 
alibaba.mozi.vds.tenant.api.service.getadmin

获取员工租户管理员信息（查询员工是否为租户管理员）
*/
type AlibabaMoziVdsTenantApiServiceGetadminAPIResponse struct {
    model.CommonResponse
    AlibabaMoziVdsTenantApiServiceGetadminResponse
}

// 获取员工租户管理员信息（查询员工是否为租户管理员） 成功返回结果
type AlibabaMoziVdsTenantApiServiceGetadminResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_vds_tenant_api_service_getadmin_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 出参
    Result   *GetEmployeeTenantAdminInfoResult `json:"result,omitempty" xml:"result,omitempty"`
}
