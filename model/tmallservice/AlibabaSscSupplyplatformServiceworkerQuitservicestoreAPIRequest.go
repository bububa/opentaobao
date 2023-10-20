package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest 工人退出网点 API请求
// alibaba.ssc.supplyplatform.serviceworker.quitservicestore
//
// 工人退出网点
type AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest struct {
	model.Params
	// 工人退出网点model
	_workerQuitServiceStoreForTopReqDto *WorkerQuitServiceStoreForTopReqDto
}

// NewAlibabaSscSupplyplatformServiceworkerQuitservicestoreRequest 初始化AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest对象
func NewAlibabaSscSupplyplatformServiceworkerQuitservicestoreRequest() *AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest {
	return &AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest) Reset() {
	r._workerQuitServiceStoreForTopReqDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.serviceworker.quitservicestore"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkerQuitServiceStoreForTopReqDto is WorkerQuitServiceStoreForTopReqDto Setter
// 工人退出网点model
func (r *AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest) SetWorkerQuitServiceStoreForTopReqDto(_workerQuitServiceStoreForTopReqDto *WorkerQuitServiceStoreForTopReqDto) error {
	r._workerQuitServiceStoreForTopReqDto = _workerQuitServiceStoreForTopReqDto
	r.Set("worker_quit_service_store_for_top_req_dto", _workerQuitServiceStoreForTopReqDto)
	return nil
}

// GetWorkerQuitServiceStoreForTopReqDto WorkerQuitServiceStoreForTopReqDto Getter
func (r AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest) GetWorkerQuitServiceStoreForTopReqDto() *WorkerQuitServiceStoreForTopReqDto {
	return r._workerQuitServiceStoreForTopReqDto
}

var poolAlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSscSupplyplatformServiceworkerQuitservicestoreRequest()
	},
}

// GetAlibabaSscSupplyplatformServiceworkerQuitservicestoreRequest 从 sync.Pool 获取 AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest
func GetAlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest() *AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest {
	return poolAlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest.Get().(*AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest)
}

// ReleaseAlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest 将 AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest 放入 sync.Pool
func ReleaseAlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest(v *AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest) {
	v.Reset()
	poolAlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest.Put(v)
}
