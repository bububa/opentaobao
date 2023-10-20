package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophybilldailyqueryAPIResponse 账单日汇总接口 API返回值
// alibaba.tcls.aelophy.bill.daily.query
//
// 账单日汇总接口
type AlibabatclsaelophybilldailyqueryAPIResponse struct {
	model.CommonResponse
	AlibabatclsaelophybilldailyqueryAPIResponseModel
}

// AlibabatclsaelophybilldailyqueryAPIResponseModel is 账单日汇总接口 成功返回结果
type AlibabatclsaelophybilldailyqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_bill_daily_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	ApiResult *ApiPageResults `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
