package mozi

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaMoziVdsTenantApiServiceDismissAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziVdsTenantApiServiceDismissAPIResponseModel).Reset()
}

// AlibabaMoziVdsTenantApiServiceDismissAPIResponseModel is MOZI解除组织主管服务 成功返回结果
type AlibabaMoziVdsTenantApiServiceDismissAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_vds_tenant_api_service_dismiss_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DismissOrganizationSupervisorResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziVdsTenantApiServiceDismissAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMoziVdsTenantApiServiceDismissAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziVdsTenantApiServiceDismissAPIResponse)
	},
}

// GetAlibabaMoziVdsTenantApiServiceDismissAPIResponse 从 sync.Pool 获取 AlibabaMoziVdsTenantApiServiceDismissAPIResponse
func GetAlibabaMoziVdsTenantApiServiceDismissAPIResponse() *AlibabaMoziVdsTenantApiServiceDismissAPIResponse {
	return poolAlibabaMoziVdsTenantApiServiceDismissAPIResponse.Get().(*AlibabaMoziVdsTenantApiServiceDismissAPIResponse)
}

// ReleaseAlibabaMoziVdsTenantApiServiceDismissAPIResponse 将 AlibabaMoziVdsTenantApiServiceDismissAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziVdsTenantApiServiceDismissAPIResponse(v *AlibabaMoziVdsTenantApiServiceDismissAPIResponse) {
	v.Reset()
	poolAlibabaMoziVdsTenantApiServiceDismissAPIResponse.Put(v)
}
