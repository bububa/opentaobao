package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商工人信息创建 APIRequest
tmall.servicecenter.worker.create

服务商工人信息创建
*/
type TmallServicecenterWorkerCreateRequest struct {
    model.Params

    // 11
    workerDto   *WorkerDto 

}

func NewTmallServicecenterWorkerCreateRequest() *TmallServicecenterWorkerCreateRequest{
    return &TmallServicecenterWorkerCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkerCreateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.worker.create"
}

func (r TmallServicecenterWorkerCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkerCreateRequest) SetWorkerDto(workerDto *WorkerDto) error {
    r.workerDto = workerDto
    r.Set("worker_dto", workerDto)
    return nil
}

func (r TmallServicecenterWorkerCreateRequest) GetWorkerDto() *WorkerDto {
    return r.workerDto
}

