package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscsupplyplatformservicecapacitysaveAPIRequest 保存服务容量 API请求
// alibaba.ssc.supplyplatform.servicecapacity.save
//
// 保存服务容量
type AlibabasscsupplyplatformservicecapacitysaveAPIRequest struct {
	model.Params
	// 目前支持两种。daily 每天容量相同；customize 定制容量，每天都不同
	_mode string
	// 目前支持两种。day 表示支持的时间粒度为天；hour 时间粒度为小时
	_timeInterval string
	// 容量数据。根据mode和time interval有不同的格式。具体格式见业务对接文档。
	_capacityData string
	// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
	_providerType string
	// 网点编码
	_providerCode string
	// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
	_providerId int64
}

// NewAlibabasscsupplyplatformservicecapacitysaveRequest 初始化AlibabasscsupplyplatformservicecapacitysaveAPIRequest对象
func NewAlibabasscsupplyplatformservicecapacitysaveRequest() *AlibabasscsupplyplatformservicecapacitysaveAPIRequest {
	return &AlibabasscsupplyplatformservicecapacitysaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasscsupplyplatformservicecapacitysaveAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.servicecapacity.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasscsupplyplatformservicecapacitysaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasscsupplyplatformservicecapacitysaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMode is Mode Setter
// 目前支持两种。daily 每天容量相同；customize 定制容量，每天都不同
func (r *AlibabasscsupplyplatformservicecapacitysaveAPIRequest) SetMode(_mode string) error {
	r._mode = _mode
	r.Set("mode", _mode)
	return nil
}

// GetMode Mode Getter
func (r AlibabasscsupplyplatformservicecapacitysaveAPIRequest) GetMode() string {
	return r._mode
}

// SetTimeInterval is TimeInterval Setter
// 目前支持两种。day 表示支持的时间粒度为天；hour 时间粒度为小时
func (r *AlibabasscsupplyplatformservicecapacitysaveAPIRequest) SetTimeInterval(_timeInterval string) error {
	r._timeInterval = _timeInterval
	r.Set("time_interval", _timeInterval)
	return nil
}

// GetTimeInterval TimeInterval Getter
func (r AlibabasscsupplyplatformservicecapacitysaveAPIRequest) GetTimeInterval() string {
	return r._timeInterval
}

// SetCapacityData is CapacityData Setter
// 容量数据。根据mode和time interval有不同的格式。具体格式见业务对接文档。
func (r *AlibabasscsupplyplatformservicecapacitysaveAPIRequest) SetCapacityData(_capacityData string) error {
	r._capacityData = _capacityData
	r.Set("capacity_data", _capacityData)
	return nil
}

// GetCapacityData CapacityData Getter
func (r AlibabasscsupplyplatformservicecapacitysaveAPIRequest) GetCapacityData() string {
	return r._capacityData
}

// SetProviderType is ProviderType Setter
// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
func (r *AlibabasscsupplyplatformservicecapacitysaveAPIRequest) SetProviderType(_providerType string) error {
	r._providerType = _providerType
	r.Set("provider_type", _providerType)
	return nil
}

// GetProviderType ProviderType Getter
func (r AlibabasscsupplyplatformservicecapacitysaveAPIRequest) GetProviderType() string {
	return r._providerType
}

// SetProviderCode is ProviderCode Setter
// 网点编码
func (r *AlibabasscsupplyplatformservicecapacitysaveAPIRequest) SetProviderCode(_providerCode string) error {
	r._providerCode = _providerCode
	r.Set("provider_code", _providerCode)
	return nil
}

// GetProviderCode ProviderCode Getter
func (r AlibabasscsupplyplatformservicecapacitysaveAPIRequest) GetProviderCode() string {
	return r._providerCode
}

// SetProviderId is ProviderId Setter
// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
func (r *AlibabasscsupplyplatformservicecapacitysaveAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r AlibabasscsupplyplatformservicecapacitysaveAPIRequest) GetProviderId() int64 {
	return r._providerId
}
