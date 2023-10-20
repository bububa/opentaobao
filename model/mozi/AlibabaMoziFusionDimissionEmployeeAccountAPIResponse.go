package mozi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziFusionDimissionEmployeeAccountAPIResponse 人员离职 API返回值
// alibaba.mozi.fusion.dimission.employee.account
//
// 人员离职并且回收账号
type AlibabaMoziFusionDimissionEmployeeAccountAPIResponse struct {
	model.CommonResponse
	AlibabaMoziFusionDimissionEmployeeAccountAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMoziFusionDimissionEmployeeAccountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziFusionDimissionEmployeeAccountAPIResponseModel).Reset()
}

// AlibabaMoziFusionDimissionEmployeeAccountAPIResponseModel is 人员离职 成功返回结果
type AlibabaMoziFusionDimissionEmployeeAccountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_fusion_dimission_employee_account_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *RemoveTenantEmployeeAndAccountResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziFusionDimissionEmployeeAccountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMoziFusionDimissionEmployeeAccountAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziFusionDimissionEmployeeAccountAPIResponse)
	},
}

// GetAlibabaMoziFusionDimissionEmployeeAccountAPIResponse 从 sync.Pool 获取 AlibabaMoziFusionDimissionEmployeeAccountAPIResponse
func GetAlibabaMoziFusionDimissionEmployeeAccountAPIResponse() *AlibabaMoziFusionDimissionEmployeeAccountAPIResponse {
	return poolAlibabaMoziFusionDimissionEmployeeAccountAPIResponse.Get().(*AlibabaMoziFusionDimissionEmployeeAccountAPIResponse)
}

// ReleaseAlibabaMoziFusionDimissionEmployeeAccountAPIResponse 将 AlibabaMoziFusionDimissionEmployeeAccountAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziFusionDimissionEmployeeAccountAPIResponse(v *AlibabaMoziFusionDimissionEmployeeAccountAPIResponse) {
	v.Reset()
	poolAlibabaMoziFusionDimissionEmployeeAccountAPIResponse.Put(v)
}
