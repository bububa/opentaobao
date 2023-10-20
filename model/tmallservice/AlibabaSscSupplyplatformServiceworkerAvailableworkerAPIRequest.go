package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest 查询可用工人 API请求
// alibaba.ssc.supplyplatform.serviceworker.availableworker
//
// 可用工人查询
type AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest struct {
	model.Params
	// 查询可用工人model
	_availableWorkerQueryForTopReqDto *AvailableWorkerQueryForTopReqDto
}

// NewAlibabaSscSupplyplatformServiceworkerAvailableworkerRequest 初始化AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest对象
func NewAlibabaSscSupplyplatformServiceworkerAvailableworkerRequest() *AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest {
	return &AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest) Reset() {
	r._availableWorkerQueryForTopReqDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.serviceworker.availableworker"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAvailableWorkerQueryForTopReqDto is AvailableWorkerQueryForTopReqDto Setter
// 查询可用工人model
func (r *AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest) SetAvailableWorkerQueryForTopReqDto(_availableWorkerQueryForTopReqDto *AvailableWorkerQueryForTopReqDto) error {
	r._availableWorkerQueryForTopReqDto = _availableWorkerQueryForTopReqDto
	r.Set("available_worker_query_for_top_req_dto", _availableWorkerQueryForTopReqDto)
	return nil
}

// GetAvailableWorkerQueryForTopReqDto AvailableWorkerQueryForTopReqDto Getter
func (r AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest) GetAvailableWorkerQueryForTopReqDto() *AvailableWorkerQueryForTopReqDto {
	return r._availableWorkerQueryForTopReqDto
}

var poolAlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSscSupplyplatformServiceworkerAvailableworkerRequest()
	},
}

// GetAlibabaSscSupplyplatformServiceworkerAvailableworkerRequest 从 sync.Pool 获取 AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest
func GetAlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest() *AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest {
	return poolAlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest.Get().(*AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest)
}

// ReleaseAlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest 将 AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest 放入 sync.Pool
func ReleaseAlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest(v *AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest) {
	v.Reset()
	poolAlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest.Put(v)
}
