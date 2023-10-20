package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugtracetoplsydlistupoutAPIResponse 零售药店查询本企业上游企业出库单据信息 API返回值
// alibaba.alihealth.drugtrace.top.lsyd.listupout
//
// 查询货主/本企业上游企业出库单据信息
type AlibabaalihealthdrugtracetoplsydlistupoutAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugtracetoplsydlistupoutAPIResponseModel
}

// AlibabaalihealthdrugtracetoplsydlistupoutAPIResponseModel is 零售药店查询本企业上游企业出库单据信息 成功返回结果
type AlibabaalihealthdrugtracetoplsydlistupoutAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_listupout_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaalihealthdrugtracetoplsydlistupoutResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
