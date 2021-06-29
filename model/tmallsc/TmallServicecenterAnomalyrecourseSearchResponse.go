package tmallsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务平台服务商一键求助单查询 APIResponse
tmall.servicecenter.anomalyrecourse.search

天猫服务平台服务商一键求助单查询
*/
type TmallServicecenterAnomalyrecourseSearchAPIResponse struct {
    model.CommonResponse
    TmallServicecenterAnomalyrecourseSearchResponse
}

type TmallServicecenterAnomalyrecourseSearchResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ResultBase `json:"result,omitempty" xml:"result,omitempty"`

    
}
