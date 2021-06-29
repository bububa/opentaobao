package mozivds

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除租户管理员服务 APIResponse
alibaba.mozi.vds.tenant.api.service.removeadmin

删除租户管理员top服务
*/
type AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse struct {
    model.CommonResponse
    AlibabaMoziVdsTenantApiServiceRemoveadminResponse
}

type AlibabaMoziVdsTenantApiServiceRemoveadminResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_vds_tenant_api_service_removeadmin_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *RemoveTenantAdminsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
