package mozi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
MOZI解除组织主管服务 APIResponse
alibaba.mozi.vds.tenant.api.service.dismiss

解除组织主管
*/
type AlibabaMoziVdsTenantApiServiceDismissAPIResponse struct {
    model.CommonResponse
    AlibabaMoziVdsTenantApiServiceDismissResponse
}

type AlibabaMoziVdsTenantApiServiceDismissResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_vds_tenant_api_service_dismiss_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *DismissOrganizationSupervisorResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
