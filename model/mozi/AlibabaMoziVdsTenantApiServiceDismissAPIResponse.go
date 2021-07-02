package mozi

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziVdsTenantApiServiceDismissAPIResponse MOZI解除组织主管服务 API返回值
// alibaba.mozi.vds.tenant.api.service.dismiss
//
// 解除组织主管
type AlibabaMoziVdsTenantApiServiceDismissAPIResponse struct {
	model.CommonResponse
	AlibabaMoziVdsTenantApiServiceDismissAPIResponseModel
}

// AlibabaMoziVdsTenantApiServiceDismissAPIResponseModel is MOZI解除组织主管服务 成功返回结果
type AlibabaMoziVdsTenantApiServiceDismissAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_vds_tenant_api_service_dismiss_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DismissOrganizationSupervisorResult `json:"result,omitempty" xml:"result,omitempty"`
}
