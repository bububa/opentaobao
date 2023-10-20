package mozi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziFusionReentryEmployeeAccountAPIResponse 重新入职并且重新启用账号 API返回值
// alibaba.mozi.fusion.reentry.employee.account
//
// 重新入职并且重新启用账号
type AlibabaMoziFusionReentryEmployeeAccountAPIResponse struct {
	model.CommonResponse
	AlibabaMoziFusionReentryEmployeeAccountAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMoziFusionReentryEmployeeAccountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziFusionReentryEmployeeAccountAPIResponseModel).Reset()
}

// AlibabaMoziFusionReentryEmployeeAccountAPIResponseModel is 重新入职并且重新启用账号 成功返回结果
type AlibabaMoziFusionReentryEmployeeAccountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_fusion_reentry_employee_account_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ReEntryTenantEmployeeAndAccountResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziFusionReentryEmployeeAccountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMoziFusionReentryEmployeeAccountAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziFusionReentryEmployeeAccountAPIResponse)
	},
}

// GetAlibabaMoziFusionReentryEmployeeAccountAPIResponse 从 sync.Pool 获取 AlibabaMoziFusionReentryEmployeeAccountAPIResponse
func GetAlibabaMoziFusionReentryEmployeeAccountAPIResponse() *AlibabaMoziFusionReentryEmployeeAccountAPIResponse {
	return poolAlibabaMoziFusionReentryEmployeeAccountAPIResponse.Get().(*AlibabaMoziFusionReentryEmployeeAccountAPIResponse)
}

// ReleaseAlibabaMoziFusionReentryEmployeeAccountAPIResponse 将 AlibabaMoziFusionReentryEmployeeAccountAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziFusionReentryEmployeeAccountAPIResponse(v *AlibabaMoziFusionReentryEmployeeAccountAPIResponse) {
	v.Reset()
	poolAlibabaMoziFusionReentryEmployeeAccountAPIResponse.Put(v)
}
