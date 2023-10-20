package tmallservice

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSscSupplyplatformServicestoreSaveAPIRequest) Reset() {
	r._serviceStoreSaveReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServicestoreSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.servicestore.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServicestoreSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSscSupplyplatformServicestoreSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaSscSupplyplatformServicestoreSaveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSscSupplyplatformServicestoreSaveRequest()
	},
}

// GetAlibabaSscSupplyplatformServicestoreSaveRequest 从 sync.Pool 获取 AlibabaSscSupplyplatformServicestoreSaveAPIRequest
func GetAlibabaSscSupplyplatformServicestoreSaveAPIRequest() *AlibabaSscSupplyplatformServicestoreSaveAPIRequest {
	return poolAlibabaSscSupplyplatformServicestoreSaveAPIRequest.Get().(*AlibabaSscSupplyplatformServicestoreSaveAPIRequest)
}

// ReleaseAlibabaSscSupplyplatformServicestoreSaveAPIRequest 将 AlibabaSscSupplyplatformServicestoreSaveAPIRequest 放入 sync.Pool
func ReleaseAlibabaSscSupplyplatformServicestoreSaveAPIRequest(v *AlibabaSscSupplyplatformServicestoreSaveAPIRequest) {
	v.Reset()
	poolAlibabaSscSupplyplatformServicestoreSaveAPIRequest.Put(v)
}
