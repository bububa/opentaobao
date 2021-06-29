package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商工人信息创建 API请求
tmall.servicecenter.worker.create

服务商工人信息创建
*/
type TmallServicecenterWorkerCreateRequest struct {
    model.Params
    // 11
    _workerDto   *WorkerDto
}

// 初始化TmallServicecenterWorkerCreateRequest对象
func NewTmallServicecenterWorkerCreateRequest() *TmallServicecenterWorkerCreateRequest{
    return &TmallServicecenterWorkerCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkerCreateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.worker.create"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkerCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkerDto Setter
// 11
func (r *TmallServicecenterWorkerCreateRequest) SetWorkerDto(_workerDto *WorkerDto) error {
    r._workerDto = _workerDto
    r.Set("worker_dto", _workerDto)
    return nil
}

// WorkerDto Getter
func (r TmallServicecenterWorkerCreateRequest) GetWorkerDto() *WorkerDto {
    return r._workerDto
}
