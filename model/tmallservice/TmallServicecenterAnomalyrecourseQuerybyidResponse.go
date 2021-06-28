package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据一键求助id查询指定服务商的一键求助单 APIResponse
tmall.servicecenter.anomalyrecourse.querybyid

根据一键求助id查询指定服务商的一键求助单
*/
type TmallServicecenterAnomalyrecourseQuerybyidAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_anomalyrecourse_querybyid_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ResultBase `json:"result,omitempty" xml:"