package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest 删除服务能力 API请求
// alibaba.ssc.supplyplatform.serviceability.delete
//
// 删除服务能力
type AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest struct {
	model.Params
	// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
	_providerType string
	// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
	_providerId int64
}

// NewAlibabaSscSupplyplatformServiceabilityDeleteRequest 初始化AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest对象
func NewAlibabaSscSupplyplatformServiceabilityDeleteRequest() *AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest {
	return &AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest) Reset() {
	r._providerType = ""
	r._providerId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.serviceability.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProviderType is ProviderType Setter
// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
func (r *AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest) SetProviderType(_providerType string) error {
	r._providerType = _providerType
	r.Set("provider_type", _providerType)
	return nil
}

// GetProviderType ProviderType Getter
func (r AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest) GetProviderType() string {
	return r._providerType
}

// SetProviderId is ProviderId Setter
// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
func (r *AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest) GetProviderId() int64 {
	return r._providerId
}

var poolAlibabaSscSupplyplatformServiceabilityDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSscSupplyplatformServiceabilityDeleteRequest()
	},
}

// GetAlibabaSscSupplyplatformServiceabilityDeleteRequest 从 sync.Pool 获取 AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest
func GetAlibabaSscSupplyplatformServiceabilityDeleteAPIRequest() *AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest {
	return poolAlibabaSscSupplyplatformServiceabilityDeleteAPIRequest.Get().(*AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest)
}

// ReleaseAlibabaSscSupplyplatformServiceabilityDeleteAPIRequest 将 AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaSscSupplyplatformServiceabilityDeleteAPIRequest(v *AlibabaSscSupplyplatformServiceabilityDeleteAPIRequest) {
	v.Reset()
	poolAlibabaSscSupplyplatformServiceabilityDeleteAPIRequest.Put(v)
}
