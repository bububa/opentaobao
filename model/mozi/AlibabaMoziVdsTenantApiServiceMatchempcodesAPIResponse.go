package mozi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
校验组-员工是否匹配 API返回值 
alibaba.mozi.vds.tenant.api.service.matchempcodes

校验组-员工是否匹配
*/
type AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse struct {
    model.CommonResponse
    AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponseModel
}

// 校验组-员工是否匹配 成功返回结果
type AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_mozi_vds_tenant_api_service_matchempcodes_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 出参
    Result   *MatchWithEmployeeResult `json:"result,omitempty" xml:"result,omitempty"`
}
