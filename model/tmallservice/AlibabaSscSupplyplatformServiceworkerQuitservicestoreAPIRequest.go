package tmallservice

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.serviceworker.quitservicestore"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
