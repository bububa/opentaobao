package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugtracetoplsydqueryrelationAPIResponse 单码关联关系查询 API返回值
// alibaba.alihealth.drugtrace.top.lsyd.query.relation
//
// 单码关联关系查询
type AlibabaalihealthdrugtracetoplsydqueryrelationAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugtracetoplsydqueryrelationAPIResponseModel
}

// AlibabaalihealthdrugtracetoplsydqueryrelationAPIResponseModel is 单码关联关系查询 成功返回结果
type AlibabaalihealthdrugtracetoplsydqueryrelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_query_relation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihealthdrugtracetoplsydqueryrelationResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
