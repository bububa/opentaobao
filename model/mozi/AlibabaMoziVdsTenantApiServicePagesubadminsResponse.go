package mozi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询租户子管理员 APIResponse
alibaba.mozi.vds.tenant.api.service.pagesubadmins

分页查询租户子管理员
*/
type AlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse struct {
    model.CommonResponse
    AlibabaMoziVdsTenantApiServicePagesubadminsResponse
}

type AlibabaMoziVdsTenantApiServicePagesubadminsResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_vds_tenant_api_service_pagesubadmins_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *PageTenantSubAdminsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
