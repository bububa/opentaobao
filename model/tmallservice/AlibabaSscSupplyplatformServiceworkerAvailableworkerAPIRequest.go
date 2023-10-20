package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscsupplyplatformserviceworkeravailableworkerAPIRequest 查询可用工人 API请求
// alibaba.ssc.supplyplatform.serviceworker.availableworker
//
// 可用工人查询
type AlibabasscsupplyplatformserviceworkeravailableworkerAPIRequest struct {
	model.Params
	// 查询可用工人model
	_availableWorkerQueryForTopReqDto *AvailableWorkerQueryForTopReqDto
}

// NewAlibabasscsupplyplatformserviceworkeravailableworkerRequest 初始化AlibabasscsupplyplatformserviceworkeravailableworkerAPIRequest对象
func NewAlibabasscsupplyplatformserviceworkeravailableworkerRequest() *AlibabasscsupplyplatformserviceworkeravailableworkerAPIRequest {
	return &AlibabasscsupplyplatformserviceworkeravailableworkerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasscsupplyplatformserviceworkeravailableworkerAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.serviceworker.availableworker"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasscsupplyplatformserviceworkeravailableworkerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasscsupplyplatformserviceworkeravailableworkerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAvailableWorkerQueryForTopReqDto is AvailableWorkerQueryForTopReqDto Setter
// 查询可用工人model
func (r *AlibabasscsupplyplatformserviceworkeravailableworkerAPIRequest) SetAvailableWorkerQueryForTopReqDto(_availableWorkerQueryForTopReqDto *AvailableWorkerQueryForTopReqDto) error {
	r._availableWorkerQueryForTopReqDto = _availableWorkerQueryForTopReqDto
	r.Set("available_worker_query_for_top_req_dto", _availableWorkerQueryForTopReqDto)
	return nil
}

// GetAvailableWorkerQueryForTopReqDto AvailableWorkerQueryForTopReqDto Getter
func (r AlibabasscsupplyplatformserviceworkeravailableworkerAPIRequest) GetAvailableWorkerQueryForTopReqDto() *AvailableWorkerQueryForTopReqDto {
	return r._availableWorkerQueryForTopReqDto
}
