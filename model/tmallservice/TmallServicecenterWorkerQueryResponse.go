package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
工人信息查询 APIResponse
tmall.servicecenter.worker.query

查询服务商对应的工人信息
*/
type TmallServicecenterWorkerQueryAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterWorkerQueryResponse `json:"tmall_servicecenter_worker_query_response,omitempty"`
}

type TmallServicecenterWorkerQueryResponse struct {

    // result
    Result   *ResultBase `json:"result,omitempty"`

}
