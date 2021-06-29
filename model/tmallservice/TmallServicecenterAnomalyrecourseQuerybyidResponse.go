package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据一键求助id查询指定服务商的一键求助单 API返回值 
tmall.servicecenter.anomalyrecourse.querybyid

根据一键求助id查询指定服务商的一键求助单
*/
type TmallServicecenterAnomalyrecourseQuerybyidAPIResponse struct {
    model.CommonResponse
    TmallServicecenterAnomalyrecourseQuerybyidResponse
}

// 根据一键求助id查询指定服务商的一键求助单 成功返回结果
type TmallServicecenterAnomalyrecourseQuerybyidResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_querybyid_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}
