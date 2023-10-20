package mozi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse 添加人员和账号复合接口 API返回值
// alibaba.mozi.fusion.addorupdate.employee.account
//
// 添加人员和账号复合接口
type AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse struct {
	model.CommonResponse
	AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponseModel).Reset()
}

// AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponseModel is 添加人员和账号复合接口 成功返回结果
type AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_fusion_addorupdate_employee_account_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AddOrUpdateTenantEmployeeAndAccountResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse)
	},
}

// GetAlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse 从 sync.Pool 获取 AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse
func GetAlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse() *AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse {
	return poolAlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse.Get().(*AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse)
}

// ReleaseAlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse 将 AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse(v *AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse) {
	v.Reset()
	poolAlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse.Put(v)
}
