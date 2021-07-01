package mozivds

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除租户管理员服务 API返回值 
alibaba.mozi.vds.tenant.api.service.removeadmin

删除租户管理员top服务
*/
type AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse struct {
    model.CommonResponse
    AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponseModel
}

// 删除租户管理员服务 成功返回结果
type AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_mozi_vds_tenant_api_service_removeadmin_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *RemoveTenantAdminsResult `json:"result,omitempty" xml:"result,omitempty"`
}
