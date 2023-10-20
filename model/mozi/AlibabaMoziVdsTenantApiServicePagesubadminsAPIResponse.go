package mozi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse 分页查询租户子管理员 API返回值
// alibaba.mozi.vds.tenant.api.service.pagesubadmins
//
// 分页查询租户子管理员
type AlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse struct {
	model.CommonResponse
	AlibabaMoziVdsTenantApiServicePagesubadminsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziVdsTenantApiServicePagesubadminsAPIResponseModel).Reset()
}

// AlibabaMoziVdsTenantApiServicePagesubadminsAPIResponseModel is 分页查询租户子管理员 成功返回结果
type AlibabaMoziVdsTenantApiServicePagesubadminsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_vds_tenant_api_service_pagesubadmins_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *PageTenantSubAdminsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziVdsTenantApiServicePagesubadminsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse)
	},
}

// GetAlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse 从 sync.Pool 获取 AlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse
func GetAlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse() *AlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse {
	return poolAlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse.Get().(*AlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse)
}

// ReleaseAlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse 将 AlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse(v *AlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse) {
	v.Reset()
	poolAlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse.Put(v)
}
