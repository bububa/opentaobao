package mozi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
按租户ID查询租户信息 APIResponse
alibaba.mozi.vds.tenant.api.service.tenantbyid

按租户ID查询租户信息
*/
type AlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse struct {
    model.CommonResponse
    AlibabaMoziVdsTenantApiServiceTenantbyidResponse
}

type AlibabaMoziVdsTenantApiServiceTenantbyidResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_vds_tenant_api_service_tenantbyid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *GetTenantByIdResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
