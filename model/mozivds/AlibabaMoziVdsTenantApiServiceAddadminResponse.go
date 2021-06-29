package mozivds

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新建租户管理员 API返回值 
alibaba.mozi.vds.tenant.api.service.addadmin

新建租户管理员
alibaba.mozi.vds.tenant.api.service.addadmin
*/
type AlibabaMoziVdsTenantApiServiceAddadminAPIResponse struct {
    model.CommonResponse
    AlibabaMoziVdsTenantApiServiceAddadminResponse
}

// 新建租户管理员 成功返回结果
type AlibabaMoziVdsTenantApiServiceAddadminResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_vds_tenant_api_service_addadmin_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *AddTenantAdminsResult `json:"result,omitempty" xml:"result,omitempty"`
}
