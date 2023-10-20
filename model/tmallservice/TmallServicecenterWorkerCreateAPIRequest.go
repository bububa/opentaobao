package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkerCreateAPIRequest 服务商工人信息创建 API请求
// tmall.servicecenter.worker.create
//
// 服务商工人信息创建
type TmallServicecenterWorkerCreateAPIRequest struct {
	model.Params
	// 11
	_workerDto *WorkerDto
}

// NewTmallServicecenterWorkerCreateRequest 初始化TmallServicecenterWorkerCreateAPIRequest对象
func NewTmallServicecenterWorkerCreateRequest() *TmallServicecenterWorkerCreateAPIRequest {
	return &TmallServicecenterWorkerCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkerCreateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.worker.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkerCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkerCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkerDto is WorkerDto Setter
// 11
func (r *TmallServicecenterWorkerCreateAPIRequest) SetWorkerDto(_workerDto *WorkerDto) error {
	r._workerDto = _workerDto
	r.Set("worker_dto", _workerDto)
	return nil
}

// GetWorkerDto WorkerDto Getter
func (r TmallServicecenterWorkerCreateAPIRequest) GetWorkerDto() *WorkerDto {
	return r._workerDto
}
