package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询工人列表 APIResponse
tmall.servicecenter.worker.querypage

服务商查询工人列表
*/
type TmallServicecenterWorkerQuerypageAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterWorkerQuerypageResponse `json:"tmall_servicecenter_worker_querypage_response,omitempty"`
}

type TmallServicecenterWorkerQuerypageResponse struct {

    // result
    Result   *ResultBase `json:"result,omitempty"`

}
