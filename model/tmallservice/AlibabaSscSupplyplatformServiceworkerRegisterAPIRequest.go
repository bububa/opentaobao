package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceworkerRegisterAPIRequest 服务商添加工人 API请求
// alibaba.ssc.supplyplatform.serviceworker.register
//
// 工人注册
type AlibabaSscSupplyplatformServiceworkerRegisterAPIRequest struct {
	model.Params
	// 工人注册model
	_workerRegisterForTopReqDto *WorkerRegisterForTopReqDto
}

// NewAlibabaSscSupplyplatformServiceworkerRegisterRequest 初始化AlibabaSscSupplyplatformServiceworkerRegisterAPIRequest对象
func NewAlibabaSscSupplyplatformServiceworkerRegisterRequest() *AlibabaSscSupplyplatformServiceworkerRegisterAPIRequest {
	return &AlibabaSscSupplyplatformServiceworkerRegisterAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSscSupplyplatformServiceworkerRegisterAPIRequest) Reset() {
	r._workerRegisterForTopReqDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServiceworkerRegisterAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.serviceworker.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServiceworkerRegisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSscSupplyplatformServiceworkerRegisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkerRegisterForTopReqDto is WorkerRegisterForTopReqDto Setter
// 工人注册model
func (r *AlibabaSscSupplyplatformServiceworkerRegisterAPIRequest) SetWorkerRegisterForTopReqDto(_workerRegisterForTopReqDto *WorkerRegisterForTopReqDto) error {
	r._workerRegisterForTopReqDto = _workerRegisterForTopReqDto
	r.Set("worker_register_for_top_req_dto", _workerRegisterForTopReqDto)
	return nil
}

// GetWorkerRegisterForTopReqDto WorkerRegisterForTopReqDto Getter
func (r AlibabaSscSupplyplatformServiceworkerRegisterAPIRequest) GetWorkerRegisterForTopReqDto() *WorkerRegisterForTopReqDto {
	return r._workerRegisterForTopReqDto
}

var poolAlibabaSscSupplyplatformServiceworkerRegisterAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSscSupplyplatformServiceworkerRegisterRequest()
	},
}

// GetAlibabaSscSupplyplatformServiceworkerRegisterRequest 从 sync.Pool 获取 AlibabaSscSupplyplatformServiceworkerRegisterAPIRequest
func GetAlibabaSscSupplyplatformServiceworkerRegisterAPIRequest() *AlibabaSscSupplyplatformServiceworkerRegisterAPIRequest {
	return poolAlibabaSscSupplyplatformServiceworkerRegisterAPIRequest.Get().(*AlibabaSscSupplyplatformServiceworkerRegisterAPIRequest)
}

// ReleaseAlibabaSscSupplyplatformServiceworkerRegisterAPIRequest 将 AlibabaSscSupplyplatformServiceworkerRegisterAPIRequest 放入 sync.Pool
func ReleaseAlibabaSscSupplyplatformServiceworkerRegisterAPIRequest(v *AlibabaSscSupplyplatformServiceworkerRegisterAPIRequest) {
	v.Reset()
	poolAlibabaSscSupplyplatformServiceworkerRegisterAPIRequest.Put(v)
}
