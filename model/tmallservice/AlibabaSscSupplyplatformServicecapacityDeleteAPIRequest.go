package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscsupplyplatformservicecapacitydeleteAPIRequest 服务容量删除 API请求
// alibaba.ssc.supplyplatform.servicecapacity.delete
//
// 服务容量删除
type AlibabasscsupplyplatformservicecapacitydeleteAPIRequest struct {
	model.Params
	// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
	_providerType string
	// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
	_providerId int64
}

// NewAlibabasscsupplyplatformservicecapacitydeleteRequest 初始化AlibabasscsupplyplatformservicecapacitydeleteAPIRequest对象
func NewAlibabasscsupplyplatformservicecapacitydeleteRequest() *AlibabasscsupplyplatformservicecapacitydeleteAPIRequest {
	return &AlibabasscsupplyplatformservicecapacitydeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasscsupplyplatformservicecapacitydeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.servicecapacity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasscsupplyplatformservicecapacitydeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasscsupplyplatformservicecapacitydeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProviderType is ProviderType Setter
// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
func (r *AlibabasscsupplyplatformservicecapacitydeleteAPIRequest) SetProviderType(_providerType string) error {
	r._providerType = _providerType
	r.Set("provider_type", _providerType)
	return nil
}

// GetProviderType ProviderType Getter
func (r AlibabasscsupplyplatformservicecapacitydeleteAPIRequest) GetProviderType() string {
	return r._providerType
}

// SetProviderId is ProviderId Setter
// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
func (r *AlibabasscsupplyplatformservicecapacitydeleteAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r AlibabasscsupplyplatformservicecapacitydeleteAPIRequest) GetProviderId() int64 {
	return r._providerId
}
