package mozi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziFusionUpdateEmployeeAccountAPIResponse 更新人员和账号属性 API返回值
// alibaba.mozi.fusion.update.employee.account
//
// 更新人员和账号基本属性
type AlibabaMoziFusionUpdateEmployeeAccountAPIResponse struct {
	model.CommonResponse
	AlibabaMoziFusionUpdateEmployeeAccountAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMoziFusionUpdateEmployeeAccountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziFusionUpdateEmployeeAccountAPIResponseModel).Reset()
}

// AlibabaMoziFusionUpdateEmployeeAccountAPIResponseModel is 更新人员和账号属性 成功返回结果
type AlibabaMoziFusionUpdateEmployeeAccountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_fusion_update_employee_account_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *UpdateTenantEmployeeAndAccountResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziFusionUpdateEmployeeAccountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMoziFusionUpdateEmployeeAccountAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziFusionUpdateEmployeeAccountAPIResponse)
	},
}

// GetAlibabaMoziFusionUpdateEmployeeAccountAPIResponse 从 sync.Pool 获取 AlibabaMoziFusionUpdateEmployeeAccountAPIResponse
func GetAlibabaMoziFusionUpdateEmployeeAccountAPIResponse() *AlibabaMoziFusionUpdateEmployeeAccountAPIResponse {
	return poolAlibabaMoziFusionUpdateEmployeeAccountAPIResponse.Get().(*AlibabaMoziFusionUpdateEmployeeAccountAPIResponse)
}

// ReleaseAlibabaMoziFusionUpdateEmployeeAccountAPIResponse 将 AlibabaMoziFusionUpdateEmployeeAccountAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziFusionUpdateEmployeeAccountAPIResponse(v *AlibabaMoziFusionUpdateEmployeeAccountAPIResponse) {
	v.Reset()
	poolAlibabaMoziFusionUpdateEmployeeAccountAPIResponse.Put(v)
}
