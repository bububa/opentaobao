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
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_worker_querycapacitytask_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // ResultBase
    
    ResultBase   *ResultBase `json:"result_base,omitempty" xml:"