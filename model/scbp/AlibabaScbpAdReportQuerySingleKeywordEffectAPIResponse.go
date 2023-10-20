package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadreportquerysinglekeywordeffectAPIResponse 单个关键词报告 API返回值
// alibaba.scbp.ad.report.query.single.keyword.effect
//
// 单个关键词报告
type AlibabascbpadreportquerysinglekeywordeffectAPIResponse struct {
	model.CommonResponse
	AlibabascbpadreportquerysinglekeywordeffectAPIResponseModel
}

// AlibabascbpadreportquerysinglekeywordeffectAPIResponseModel is 单个关键词报告 成功返回结果
type AlibabascbpadreportquerysinglekeywordeffectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_report_query_single_keyword_effect_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *KeywordReportDto `json:"result,omitempty" xml:"result,omitempty"`
}
