package mozi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse 按租户ID查询租户信息 API返回值
// alibaba.mozi.vds.tenant.api.service.tenantbyid
//
// 按租户ID查询租户信息
type AlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse struct {
	model.CommonResponse
	AlibabaMoziVdsTenantApiServiceTenantbyidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziVdsTenantApiServiceTenantbyidAPIResponseModel).Reset()
}

// AlibabaMoziVdsTenantApiServiceTenantbyidAPIResponseModel is 按租户ID查询租户信息 成功返回结果
type AlibabaMoziVdsTenantApiServiceTenantbyidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_vds_tenant_api_service_tenantbyid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *GetTenantByIdResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziVdsTenantApiServiceTenantbyidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse)
	},
}

// GetAlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse 从 sync.Pool 获取 AlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse
func GetAlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse() *AlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse {
	return poolAlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse.Get().(*AlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse)
}

// ReleaseAlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse 将 AlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse(v *AlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse) {
	v.Reset()
	poolAlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse.Put(v)
}
