package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscsupplyplatformservicestoresaveAPIRequest 保存网点 API请求
// alibaba.ssc.supplyplatform.servicestore.save
//
// 网点创建、修改
type AlibabasscsupplyplatformservicestoresaveAPIRequest struct {
	model.Params
	// 入参
	_serviceStoreSaveReq *ServiceStoreSaveForTopReqDto
}

// NewAlibabasscsupplyplatformservicestoresaveRequest 初始化AlibabasscsupplyplatformservicestoresaveAPIRequest对象
func NewAlibabasscsupplyplatformservicestoresaveRequest() *AlibabasscsupplyplatformservicestoresaveAPIRequest {
	return &AlibabasscsupplyplatformservicestoresaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasscsupplyplatformservicestoresaveAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.servicestore.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasscsupplyplatformservicestoresaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasscsupplyplatformservicestoresaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceStoreSaveReq is ServiceStoreSaveReq Setter
// 入参
func (r *AlibabasscsupplyplatformservicestoresaveAPIRequest) SetServiceStoreSaveReq(_serviceStoreSaveReq *ServiceStoreSaveForTopReqDto) error {
	r._serviceStoreSaveReq = _serviceStoreSaveReq
	r.Set("service_store_save_req", _serviceStoreSaveReq)
	return nil
}

// GetServiceStoreSaveReq ServiceStoreSaveReq Getter
func (r AlibabasscsupplyplatformservicestoresaveAPIRequest) GetServiceStoreSaveReq() *ServiceStoreSaveForTopReqDto {
	return r._serviceStoreSaveReq
}
