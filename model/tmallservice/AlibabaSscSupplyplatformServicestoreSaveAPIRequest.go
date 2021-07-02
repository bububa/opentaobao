package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServicestoreSaveAPIRequest 保存网点 API请求
// alibaba.ssc.supplyplatform.servicestore.save
//
// 网点创建、修改
type AlibabaSscSupplyplatformServicestoreSaveAPIRequest struct {
	model.Params
	// 入参
	_serviceStoreSaveReq *ServiceStoreSaveForTopReqDto
}

// NewAlibabaSscSupplyplatformServicestoreSaveRequest 初始化AlibabaSscSupplyplatformServicestoreSaveAPIRequest对象
func NewAlibabaSscSupplyplatformServicestoreSaveRequest() *AlibabaSscSupplyplatformServicestoreSaveAPIRequest {
	return &AlibabaSscSupplyplatformServicestoreSaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServicestoreSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.servicestore.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServicestoreSaveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetServiceStoreSaveReq is ServiceStoreSaveReq Setter
// 入参
func (r *AlibabaSscSupplyplatformServicestoreSaveAPIRequest) SetServiceStoreSaveReq(_serviceStoreSaveReq *ServiceStoreSaveForTopReqDto) error {
	r._serviceStoreSaveReq = _serviceStoreSaveReq
	r.Set("service_store_save_req", _serviceStoreSaveReq)
	return nil
}

// GetServiceStoreSaveReq ServiceStoreSaveReq Getter
func (r AlibabaSscSupplyplatformServicestoreSaveAPIRequest) GetServiceStoreSaveReq() *ServiceStoreSaveForTopReqDto {
	return r._serviceStoreSaveReq
}
