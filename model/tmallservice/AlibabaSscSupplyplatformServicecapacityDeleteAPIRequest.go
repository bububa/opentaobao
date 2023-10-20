package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest 服务容量删除 API请求
// alibaba.ssc.supplyplatform.servicecapacity.delete
//
// 服务容量删除
type AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest struct {
	model.Params
	// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
	_providerType string
	// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
	_providerId int64
}

// NewAlibabaSscSupplyplatformServicecapacityDeleteRequest 初始化AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest对象
func NewAlibabaSscSupplyplatformServicecapacityDeleteRequest() *AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest {
	return &AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest) Reset() {
	r._providerType = ""
	r._providerId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.servicecapacity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProviderType is ProviderType Setter
// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
func (r *AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest) SetProviderType(_providerType string) error {
	r._providerType = _providerType
	r.Set("provider_type", _providerType)
	return nil
}

// GetProviderType ProviderType Getter
func (r AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest) GetProviderType() string {
	return r._providerType
}

// SetProviderId is ProviderId Setter
// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
func (r *AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest) GetProviderId() int64 {
	return r._providerId
}

var poolAlibabaSscSupplyplatformServicecapacityDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSscSupplyplatformServicecapacityDeleteRequest()
	},
}

// GetAlibabaSscSupplyplatformServicecapacityDeleteRequest 从 sync.Pool 获取 AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest
func GetAlibabaSscSupplyplatformServicecapacityDeleteAPIRequest() *AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest {
	return poolAlibabaSscSupplyplatformServicecapacityDeleteAPIRequest.Get().(*AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest)
}

// ReleaseAlibabaSscSupplyplatformServicecapacityDeleteAPIRequest 将 AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaSscSupplyplatformServicecapacityDeleteAPIRequest(v *AlibabaSscSupplyplatformServicecapacityDeleteAPIRequest) {
	v.Reset()
	poolAlibabaSscSupplyplatformServicecapacityDeleteAPIRequest.Put(v)
}
