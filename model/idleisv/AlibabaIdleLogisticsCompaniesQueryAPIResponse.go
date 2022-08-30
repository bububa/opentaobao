package idleisv

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleLogisticsCompaniesQueryAPIResponse 快递公司列表查询 API返回值
// alibaba.idle.logistics.companies.query
//
// 支持发货的快递公司列表查询
type AlibabaIdleLogisticsCompaniesQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIdleLogisticsCompaniesQueryAPIResponseModel
}

// AlibabaIdleLogisticsCompaniesQueryAPIResponseModel is 快递公司列表查询 成功返回结果
type AlibabaIdleLogisticsCompaniesQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_logistics_companies_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultErrCode string `json:"result_err_code,omitempty" xml:"result_err_code,omitempty"`
	// 错误信息
	ResultErrMsg string `json:"result_err_msg,omitempty" xml:"result_err_msg,omitempty"`
	// 快递公司列表返回参数
	LogisticsRespResult *LogisticsRespResult `json:"logistics_resp_result,omitempty" xml:"logistics_resp_result,omitempty"`
	// 调用结果
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}
