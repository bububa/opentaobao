package mozi

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamozifusionupdateemployeeaccountAPIResponse 更新人员和账号属性 API返回值
// alibaba.mozi.fusion.update.employee.account
//
// 更新人员和账号基本属性
type AlibabamozifusionupdateemployeeaccountAPIResponse struct {
	model.CommonResponse
	AlibabamozifusionupdateemployeeaccountAPIResponseModel
}

// AlibabamozifusionupdateemployeeaccountAPIResponseModel is 更新人员和账号属性 成功返回结果
type AlibabamozifusionupdateemployeeaccountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_fusion_update_employee_account_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *UpdateTenantEmployeeAndAccountResult `json:"result,omitempty" xml:"result,omitempty"`
}
