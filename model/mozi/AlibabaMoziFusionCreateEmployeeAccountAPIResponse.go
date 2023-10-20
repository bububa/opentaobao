package mozi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziFusionCreateEmployeeAccountAPIResponse 创建MOZI自建人员和账号 API返回值
// alibaba.mozi.fusion.create.employee.account
//
// 创建MOZI自建人员和账号
type AlibabaMoziFusionCreateEmployeeAccountAPIResponse struct {
	model.CommonResponse
	AlibabaMoziFusionCreateEmployeeAccountAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMoziFusionCreateEmployeeAccountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziFusionCreateEmployeeAccountAPIResponseModel).Reset()
}

// AlibabaMoziFusionCreateEmployeeAccountAPIResponseModel is 创建MOZI自建人员和账号 成功返回结果
type AlibabaMoziFusionCreateEmployeeAccountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_fusion_create_employee_account_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CreateTenantEmployeeAndAccountResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziFusionCreateEmployeeAccountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMoziFusionCreateEmployeeAccountAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziFusionCreateEmployeeAccountAPIResponse)
	},
}

// GetAlibabaMoziFusionCreateEmployeeAccountAPIResponse 从 sync.Pool 获取 AlibabaMoziFusionCreateEmployeeAccountAPIResponse
func GetAlibabaMoziFusionCreateEmployeeAccountAPIResponse() *AlibabaMoziFusionCreateEmployeeAccountAPIResponse {
	return poolAlibabaMoziFusionCreateEmployeeAccountAPIResponse.Get().(*AlibabaMoziFusionCreateEmployeeAccountAPIResponse)
}

// ReleaseAlibabaMoziFusionCreateEmployeeAccountAPIResponse 将 AlibabaMoziFusionCreateEmployeeAccountAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziFusionCreateEmployeeAccountAPIResponse(v *AlibabaMoziFusionCreateEmployeeAccountAPIResponse) {
	v.Reset()
	poolAlibabaMoziFusionCreateEmployeeAccountAPIResponse.Put(v)
}
