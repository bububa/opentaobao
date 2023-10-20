package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugtracetopyljgquerylistpartsAPIResponse 往来单位查询 API返回值
// alibaba.alihealth.drugtrace.top.yljg.query.listparts
//
// 查询往来单位列表
type AlibabaalihealthdrugtracetopyljgquerylistpartsAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugtracetopyljgquerylistpartsAPIResponseModel
}

// AlibabaalihealthdrugtracetopyljgquerylistpartsAPIResponseModel is 往来单位查询 成功返回结果
type AlibabaalihealthdrugtracetopyljgquerylistpartsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_query_listparts_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaalihealthdrugtracetopyljgquerylistpartsResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
