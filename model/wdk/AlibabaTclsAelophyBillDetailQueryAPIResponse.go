package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophybilldetailqueryAPIResponse 账单明细接口 API返回值
// alibaba.tcls.aelophy.bill.detail.query
//
// 账单明细接口
type AlibabatclsaelophybilldetailqueryAPIResponse struct {
	model.CommonResponse
	AlibabatclsaelophybilldetailqueryAPIResponseModel
}

// AlibabatclsaelophybilldetailqueryAPIResponseModel is 账单明细接口 成功返回结果
type AlibabatclsaelophybilldetailqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_bill_detail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	ApiResult *ApiPageResults `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
