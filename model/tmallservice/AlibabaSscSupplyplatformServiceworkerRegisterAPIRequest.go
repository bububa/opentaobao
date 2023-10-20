package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscsupplyplatformserviceworkerregisterAPIRequest 服务商添加工人 API请求
// alibaba.ssc.supplyplatform.serviceworker.register
//
// 工人注册
type AlibabasscsupplyplatformserviceworkerregisterAPIRequest struct {
	model.Params
	// 工人注册model
	_workerRegisterForTopReqDto *WorkerRegisterForTopReqDto
}

// NewAlibabasscsupplyplatformserviceworkerregisterRequest 初始化AlibabasscsupplyplatformserviceworkerregisterAPIRequest对象
func NewAlibabasscsupplyplatformserviceworkerregisterRequest() *AlibabasscsupplyplatformserviceworkerregisterAPIRequest {
	return &AlibabasscsupplyplatformserviceworkerregisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasscsupplyplatformserviceworkerregisterAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.serviceworker.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasscsupplyplatformserviceworkerregisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasscsupplyplatformserviceworkerregisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkerRegisterForTopReqDto is WorkerRegisterForTopReqDto Setter
// 工人注册model
func (r *AlibabasscsupplyplatformserviceworkerregisterAPIRequest) SetWorkerRegisterForTopReqDto(_workerRegisterForTopReqDto *WorkerRegisterForTopReqDto) error {
	r._workerRegisterForTopReqDto = _workerRegisterForTopReqDto
	r.Set("worker_register_for_top_req_dto", _workerRegisterForTopReqDto)
	return nil
}

// GetWorkerRegisterForTopReqDto WorkerRegisterForTopReqDto Getter
func (r AlibabasscsupplyplatformserviceworkerregisterAPIRequest) GetWorkerRegisterForTopReqDto() *WorkerRegisterForTopReqDto {
	return r._workerRegisterForTopReqDto
}
