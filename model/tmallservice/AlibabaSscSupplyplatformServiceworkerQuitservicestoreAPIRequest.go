package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscsupplyplatformserviceworkerquitservicestoreAPIRequest 工人退出网点 API请求
// alibaba.ssc.supplyplatform.serviceworker.quitservicestore
//
// 工人退出网点
type AlibabasscsupplyplatformserviceworkerquitservicestoreAPIRequest struct {
	model.Params
	// 工人退出网点model
	_workerQuitServiceStoreForTopReqDto *WorkerQuitServiceStoreForTopReqDto
}

// NewAlibabasscsupplyplatformserviceworkerquitservicestoreRequest 初始化AlibabasscsupplyplatformserviceworkerquitservicestoreAPIRequest对象
func NewAlibabasscsupplyplatformserviceworkerquitservicestoreRequest() *AlibabasscsupplyplatformserviceworkerquitservicestoreAPIRequest {
	return &AlibabasscsupplyplatformserviceworkerquitservicestoreAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasscsupplyplatformserviceworkerquitservicestoreAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.serviceworker.quitservicestore"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasscsupplyplatformserviceworkerquitservicestoreAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasscsupplyplatformserviceworkerquitservicestoreAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkerQuitServiceStoreForTopReqDto is WorkerQuitServiceStoreForTopReqDto Setter
// 工人退出网点model
func (r *AlibabasscsupplyplatformserviceworkerquitservicestoreAPIRequest) SetWorkerQuitServiceStoreForTopReqDto(_workerQuitServiceStoreForTopReqDto *WorkerQuitServiceStoreForTopReqDto) error {
	r._workerQuitServiceStoreForTopReqDto = _workerQuitServiceStoreForTopReqDto
	r.Set("worker_quit_service_store_for_top_req_dto", _workerQuitServiceStoreForTopReqDto)
	return nil
}

// GetWorkerQuitServiceStoreForTopReqDto WorkerQuitServiceStoreForTopReqDto Getter
func (r AlibabasscsupplyplatformserviceworkerquitservicestoreAPIRequest) GetWorkerQuitServiceStoreForTopReqDto() *WorkerQuitServiceStoreForTopReqDto {
	return r._workerQuitServiceStoreForTopReqDto
}
