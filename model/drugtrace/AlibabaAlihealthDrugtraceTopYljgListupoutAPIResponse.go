package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugtracetopyljglistupoutAPIResponse 医疗机构查询本企业上游企业出库单据信息 API返回值
// alibaba.alihealth.drugtrace.top.yljg.listupout
//
// 查询货主/本企业上游企业出库单据信息
type AlibabaalihealthdrugtracetopyljglistupoutAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugtracetopyljglistupoutAPIResponseModel
}

// AlibabaalihealthdrugtracetopyljglistupoutAPIResponseModel is 医疗机构查询本企业上游企业出库单据信息 成功返回结果
type AlibabaalihealthdrugtracetopyljglistupoutAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_listupout_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaalihealthdrugtracetopyljglistupoutResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
