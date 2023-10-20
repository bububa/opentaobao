package mozivds

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse 删除租户管理员服务 API返回值
// alibaba.mozi.vds.tenant.api.service.removeadmin
//
// 删除租户管理员top服务
type AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse struct {
	model.CommonResponse
	AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponseModel).Reset()
}

// AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponseModel is 删除租户管理员服务 成功返回结果
type AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_vds_tenant_api_service_removeadmin_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *RemoveTenantAdminsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse)
	},
}

// GetAlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse 从 sync.Pool 获取 AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse
func GetAlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse() *AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse {
	return poolAlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse.Get().(*AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse)
}

// ReleaseAlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse 将 AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse(v *AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse) {
	v.Reset()
	poolAlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse.Put(v)
}
