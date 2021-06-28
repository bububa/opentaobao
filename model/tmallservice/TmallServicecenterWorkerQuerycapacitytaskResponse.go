package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询需求容量 APIResponse
tmall.servicecenter.worker.querycapacitytask

查询需求容量
*/
type TmallServicecenterWorkerQuerycapacitytaskAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkerQuerycapacitytaskResponse
}

type TmallServicecenterWorkerQuerycapacitytaskResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_worker_querycapacitytask_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // ResultBase
    
    ResultBase   *ResultBase `json:"result_base,omitempty" xml:"result_base,omitempty"`

    
}
