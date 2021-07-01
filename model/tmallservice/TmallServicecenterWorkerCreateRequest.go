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
type TmallServicecenterWorkerCreateAPIRequest struct {
    model.Params
    // 11
    _workerDto   *WorkerDTO
}

// 初始化TmallServicecenterWorkerCreateAPIRequest对象
func NewTmallServicecenterWorkerCreateRequest() *TmallServicecenterWorkerCreateAPIRequest{
    return &TmallServicecenterWorkerCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkerCreateAPIRequest) GetApiMethodName() string {
    return "tmall.servicecenter.worker.create"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkerCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkerDto Setter
// 11
func (r *TmallServicecenterWorkerCreateAPIRequest) SetWorkerDto(_workerDto *WorkerDTO) error {
    r._workerDto = _workerDto
    r.Set("worker_dto", _workerDto)
    return nil
}

// WorkerDto Getter
func (r TmallServicecenterWorkerCreateAPIRequest) GetWorkerDto() *WorkerDTO {
    return r._workerDto
}
