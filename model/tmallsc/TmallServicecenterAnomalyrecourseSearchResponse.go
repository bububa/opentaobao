package tmallsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务平台服务商一键求助单查询 API返回值 
tmall.servicecenter.anomalyrecourse.search

天猫服务平台服务商一键求助单查询
*/
type TmallServicecenterAnomalyrecourseSearchAPIResponse struct {
    model.CommonResponse
    TmallServicecenterAnomalyrecourseSearchResponse
}

// 天猫服务平台服务商一键求助单查询 成功返回结果
type TmallServicecenterAnomalyrecourseSearchResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}
