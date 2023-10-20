package mozi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse 校验组-员工是否匹配 API返回值
// alibaba.mozi.vds.tenant.api.service.matchempcodes
//
// 校验组-员工是否匹配
type AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse struct {
	model.CommonResponse
	AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponseModel).Reset()
}

// AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponseModel is 校验组-员工是否匹配 成功返回结果
type AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_vds_tenant_api_service_matchempcodes_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *MatchWithEmployeeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse)
	},
}

// GetAlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse 从 sync.Pool 获取 AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse
func GetAlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse() *AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse {
	return poolAlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse.Get().(*AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse)
}

// ReleaseAlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse 将 AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse(v *AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse) {
	v.Reset()
	poolAlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse.Put(v)
}
