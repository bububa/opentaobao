package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据一键求助id查询指定服务商的一键求助单 APIResponse
tmall.servicecenter.anomalyrecourse.querybyid

根据一键求助id查询指定服务商的一键求助单
*/
type TmallServicecenterAnomalyrecourseQuerybyidAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterAnomalyrecourseQuerybyidResponse `json:"tmall_servicecenter_anomalyrecourse_querybyid_response,omitempty"`
}

type TmallServicecenterAnomalyrecourseQuerybyidResponse struct {

    // result
    Result   *ResultBase `json:"result,omitempty"`

}
