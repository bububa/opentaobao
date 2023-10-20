package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeadvisersubmitaccountAPIResponse 顾问入职离职 API返回值
// alibaba.alihouse.newhome.adviser.submit.account
//
// 顾问入职离职
type AlibabaalihousenewhomeadvisersubmitaccountAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeadvisersubmitaccountAPIResponseModel
}

// AlibabaalihousenewhomeadvisersubmitaccountAPIResponseModel is 顾问入职离职 成功返回结果
type AlibabaalihousenewhomeadvisersubmitaccountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_adviser_submit_account_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomeadvisersubmitaccountResult `json:"result,omitempty" xml:"result,omitempty"`
}
