package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkercreateAPIRequest 服务商工人信息创建 API请求
// tmall.servicecenter.worker.create
//
// 服务商工人信息创建
type TmallservicecenterworkercreateAPIRequest struct {
	model.Params
	// 11
	_workerDto *WorkerDto
}

// NewTmallservicecenterworkercreateRequest 初始化TmallservicecenterworkercreateAPIRequest对象
func NewTmallservicecenterworkercreateRequest() *TmallservicecenterworkercreateAPIRequest {
	return &TmallservicecenterworkercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkercreateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.worker.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkerDto is WorkerDto Setter
// 11
func (r *TmallservicecenterworkercreateAPIRequest) SetWorkerDto(_workerDto *WorkerDto) error {
	r._workerDto = _workerDto
	r.Set("worker_dto", _workerDto)
	return nil
}

// GetWorkerDto WorkerDto Getter
func (r TmallservicecenterworkercreateAPIRequest) GetWorkerDto() *WorkerDto {
	return r._workerDto
}
