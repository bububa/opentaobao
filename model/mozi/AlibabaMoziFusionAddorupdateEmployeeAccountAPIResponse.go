package mozi

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamozifusionaddorupdateemployeeaccountAPIResponse 添加人员和账号复合接口 API返回值
// alibaba.mozi.fusion.addorupdate.employee.account
//
// 添加人员和账号复合接口
type AlibabamozifusionaddorupdateemployeeaccountAPIResponse struct {
	model.CommonResponse
	AlibabamozifusionaddorupdateemployeeaccountAPIResponseModel
}

// AlibabamozifusionaddorupdateemployeeaccountAPIResponseModel is 添加人员和账号复合接口 成功返回结果
type AlibabamozifusionaddorupdateemployeeaccountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_fusion_addorupdate_employee_account_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AddOrUpdateTenantEmployeeAndAccountResult `json:"result,omitempty" xml:"result,omitempty"`
}
