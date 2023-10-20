package mozivds

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziVdsTenantApiServiceAddadminAPIResponse 新建租户管理员 API返回值
// alibaba.mozi.vds.tenant.api.service.addadmin
//
// 新建租户管理员
// alibaba.mozi.vds.tenant.api.service.addadmin
type AlibabaMoziVdsTenantApiServiceAddadminAPIResponse struct {
	model.CommonResponse
	AlibabaMoziVdsTenantApiServiceAddadminAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMoziVdsTenantApiServiceAddadminAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziVdsTenantApiServiceAddadminAPIResponseModel).Reset()
}

// AlibabaMoziVdsTenantApiServiceAddadminAPIResponseModel is 新建租户管理员 成功返回结果
type AlibabaMoziVdsTenantApiServiceAddadminAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_vds_tenant_api_service_addadmin_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AddTenantAdminsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziVdsTenantApiServiceAddadminAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMoziVdsTenantApiServiceAddadminAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziVdsTenantApiServiceAddadminAPIResponse)
	},
}

// GetAlibabaMoziVdsTenantApiServiceAddadminAPIResponse 从 sync.Pool 获取 AlibabaMoziVdsTenantApiServiceAddadminAPIResponse
func GetAlibabaMoziVdsTenantApiServiceAddadminAPIResponse() *AlibabaMoziVdsTenantApiServiceAddadminAPIResponse {
	return poolAlibabaMoziVdsTenantApiServiceAddadminAPIResponse.Get().(*AlibabaMoziVdsTenantApiServiceAddadminAPIResponse)
}

// ReleaseAlibabaMoziVdsTenantApiServiceAddadminAPIResponse 将 AlibabaMoziVdsTenantApiServiceAddadminAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziVdsTenantApiServiceAddadminAPIResponse(v *AlibabaMoziVdsTenantApiServiceAddadminAPIResponse) {
	v.Reset()
	poolAlibabaMoziVdsTenantApiServiceAddadminAPIResponse.Put(v)
}
