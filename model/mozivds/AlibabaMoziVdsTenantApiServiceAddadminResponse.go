package mozivds

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新建租户管理员 APIResponse
alibaba.mozi.vds.tenant.api.service.addadmin

新建租户管理员
alibaba.mozi.vds.tenant.api.service.addadmin
*/
type AlibabaMoziVdsTenantApiServiceAddadminAPIResponse struct {
    model.CommonResponse
    AlibabaMoziVdsTenantApiServiceAddadminResponse
}

type AlibabaMoziVdsTenantApiServiceAddadminResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_vds_tenant_api_service_addadmin_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *AddTenantAdminsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
