package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugtracetopyljgqueryupbillcodeAPIResponse 通过一个码，查询这个码对应的上游企业出库单的单据号 API返回值
// alibaba.alihealth.drugtrace.top.yljg.query.upbillcode
//
// 一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号
type AlibabaalihealthdrugtracetopyljgqueryupbillcodeAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugtracetopyljgqueryupbillcodeAPIResponseModel
}

// AlibabaalihealthdrugtracetopyljgqueryupbillcodeAPIResponseModel is 通过一个码，查询这个码对应的上游企业出库单的单据号 成功返回结果
type AlibabaalihealthdrugtracetopyljgqueryupbillcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_query_upbillcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihealthdrugtracetopyljgqueryupbillcodeResult `json:"result,omitempty" xml:"result,omitempty"`
}
