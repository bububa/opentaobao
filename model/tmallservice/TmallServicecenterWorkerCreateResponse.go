package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务商工人信息创建 APIResponse
tmall.servicecenter.worker.create

服务商工人信息创建
*/
type TmallServicecenterWorkerCreateAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterWorkerCreateResponse `json:"tmall_servicecenter_worker_create_response,omitempty"`
}

type TmallServicecenterWorkerCreateResponse struct {

    // result
    Result   *ResultBase `json:"result,omitempty"`

}
