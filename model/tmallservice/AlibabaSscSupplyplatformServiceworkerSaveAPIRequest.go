package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscsupplyplatformserviceworkersaveAPIRequest 服务商绑定工人 API请求
// alibaba.ssc.supplyplatform.serviceworker.save
//
// 服务商将上传工人与服务商自己建立关系，需要将工人的服务区域和住址回传
type AlibabasscsupplyplatformserviceworkersaveAPIRequest struct {
	model.Params
	// 工人保存参数
	_workerSaveForTopReqDto *WorkerSaveForTopReqDto
}

// NewAlibabasscsupplyplatformserviceworkersaveRequest 初始化AlibabasscsupplyplatformserviceworkersaveAPIRequest对象
func NewAlibabasscsupplyplatformserviceworkersaveRequest() *AlibabasscsupplyplatformserviceworkersaveAPIRequest {
	return &AlibabasscsupplyplatformserviceworkersaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasscsupplyplatformserviceworkersaveAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.serviceworker.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasscsupplyplatformserviceworkersaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasscsupplyplatformserviceworkersaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkerSaveForTopReqDto is WorkerSaveForTopReqDto Setter
// 工人保存参数
func (r *AlibabasscsupplyplatformserviceworkersaveAPIRequest) SetWorkerSaveForTopReqDto(_workerSaveForTopReqDto *WorkerSaveForTopReqDto) error {
	r._workerSaveForTopReqDto = _workerSaveForTopReqDto
	r.Set("worker_save_for_top_req_dto", _workerSaveForTopReqDto)
	return nil
}

// GetWorkerSaveForTopReqDto WorkerSaveForTopReqDto Getter
func (r AlibabasscsupplyplatformserviceworkersaveAPIRequest) GetWorkerSaveForTopReqDto() *WorkerSaveForTopReqDto {
	return r._workerSaveForTopReqDto
}
