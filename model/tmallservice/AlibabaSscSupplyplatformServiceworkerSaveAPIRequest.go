package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceworkerSaveAPIRequest 服务商绑定工人 API请求
// alibaba.ssc.supplyplatform.serviceworker.save
//
// 服务商将上传工人与服务商自己建立关系，需要将工人的服务区域和住址回传
type AlibabaSscSupplyplatformServiceworkerSaveAPIRequest struct {
	model.Params
	// 工人保存参数
	_workerSaveForTopReqDto *WorkerSaveForTopReqDto
}

// NewAlibabaSscSupplyplatformServiceworkerSaveRequest 初始化AlibabaSscSupplyplatformServiceworkerSaveAPIRequest对象
func NewAlibabaSscSupplyplatformServiceworkerSaveRequest() *AlibabaSscSupplyplatformServiceworkerSaveAPIRequest {
	return &AlibabaSscSupplyplatformServiceworkerSaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServiceworkerSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.serviceworker.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServiceworkerSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSscSupplyplatformServiceworkerSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkerSaveForTopReqDto is WorkerSaveForTopReqDto Setter
// 工人保存参数
func (r *AlibabaSscSupplyplatformServiceworkerSaveAPIRequest) SetWorkerSaveForTopReqDto(_workerSaveForTopReqDto *WorkerSaveForTopReqDto) error {
	r._workerSaveForTopReqDto = _workerSaveForTopReqDto
	r.Set("worker_save_for_top_req_dto", _workerSaveForTopReqDto)
	return nil
}

// GetWorkerSaveForTopReqDto WorkerSaveForTopReqDto Getter
func (r AlibabaSscSupplyplatformServiceworkerSaveAPIRequest) GetWorkerSaveForTopReqDto() *WorkerSaveForTopReqDto {
	return r._workerSaveForTopReqDto
}
