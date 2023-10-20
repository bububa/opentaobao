package mozi

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamozifusionreentryemployeeaccountAPIResponse 重新入职并且重新启用账号 API返回值
// alibaba.mozi.fusion.reentry.employee.account
//
// 重新入职并且重新启用账号
type AlibabamozifusionreentryemployeeaccountAPIResponse struct {
	model.CommonResponse
	AlibabamozifusionreentryemployeeaccountAPIResponseModel
}

// AlibabamozifusionreentryemployeeaccountAPIResponseModel is 重新入职并且重新启用账号 成功返回结果
type AlibabamozifusionreentryemployeeaccountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_fusion_reentry_employee_account_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ReEntryTenantEmployeeAndAccountResult `json:"result,omitempty" xml:"result,omitempty"`
}
