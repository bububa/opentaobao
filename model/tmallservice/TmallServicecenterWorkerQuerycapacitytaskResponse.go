package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询需求容量 APIResponse
tmall.servicecenter.worker.querycapacitytask

查询需求容量
*/
type TmallServicecenterWorkerQuerycapacitytaskAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterWorkerQuerycapacitytaskResponse `json:"tmall_servicecenter_worker_querycapacitytask_response,omitempty"`
}

type TmallServicecenterWorkerQuerycapacitytaskResponse struct {

    // ResultBase
    ResultBase   *ResultBase `json:"result_base,omitempty"`

}
