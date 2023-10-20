package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpaccountbudgetgetAPIResponse 查询日消耗预算 API返回值
// alibaba.scbp.account.budget.get
//
// 查询日消耗预算
type AlibabascbpaccountbudgetgetAPIResponse struct {
	model.CommonResponse
	AlibabascbpaccountbudgetgetAPIResponseModel
}

// AlibabascbpaccountbudgetgetAPIResponseModel is 查询日消耗预算 成功返回结果
type AlibabascbpaccountbudgetgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_account_budget_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回今日预算
	Budget string `json:"budget,omitempty" xml:"budget,omitempty"`
}
