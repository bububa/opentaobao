package tmallsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务平台服务商一键求助单查询 APIResponse
tmall.servicecenter.anomalyrecourse.search

天猫服务平台服务商一键求助单查询
*/
type TmallServicecenterAnomalyrecourseSearchAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterAnomalyrecourseSearchResponse `json:"tmall_servicecenter_anomalyrecourse_search_response,omitempty"` 
    TmallServicecenterAnomalyrecourseSearchResponse
}

/* model for simplify = false
type TmallServicecenterAnomalyrecourseSearchResponse struct {

    // result
    
    Result  *struct {
        ResultBase  *ResultBase `json:"result_base,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallServicecenterAnomalyrecourseSearchResponse struct {

    // result
    Result   *ResultBase `json:"result,omitempty"`

}
