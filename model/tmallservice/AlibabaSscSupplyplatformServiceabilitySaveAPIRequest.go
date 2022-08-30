package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceabilitySaveAPIRequest 保存服务能力 API请求
// alibaba.ssc.supplyplatform.serviceability.save
//
// 保存服务能力
type AlibabaSscSupplyplatformServiceabilitySaveAPIRequest struct {
	model.Params
	// 目前包含三种。in_store 到店；at_home 上门；transmit_service 寄修。请根据实际支持的履约类型填写
	_fulfilTypeList []string
	// 服务sku，具体的sku列表可以从服务商工作台的类目树获取
	_serviceSkuCodeList []string
	// 菜鸟地址编码，各级地址均可（全国、省、市、区、街道），根据实际支持的地区填写。当支持的履约类型包含上门时，必填
	_areaCodeList []int64
	// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
	_providerType string
	// 服务提供者code。外部与天猫对接时，自己定义的code用这个字段
	_providerCode string
	// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
	_providerId int64
}

// NewAlibabaSscSupplyplatformServiceabilitySaveRequest 初始化AlibabaSscSupplyplatformServiceabilitySaveAPIRequest对象
func NewAlibabaSscSupplyplatformServiceabilitySaveRequest() *AlibabaSscSupplyplatformServiceabilitySaveAPIRequest {
	return &AlibabaSscSupplyplatformServiceabilitySaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServiceabilitySaveAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.serviceability.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServiceabilitySaveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFulfilTypeList is FulfilTypeList Setter
// 目前包含三种。in_store 到店；at_home 上门；transmit_service 寄修。请根据实际支持的履约类型填写
func (r *AlibabaSscSupplyplatformServiceabilitySaveAPIRequest) SetFulfilTypeList(_fulfilTypeList []string) error {
	r._fulfilTypeList = _fulfilTypeList
	r.Set("fulfil_type_list", _fulfilTypeList)
	return nil
}

// GetFulfilTypeList FulfilTypeList Getter
func (r AlibabaSscSupplyplatformServiceabilitySaveAPIRequest) GetFulfilTypeList() []string {
	return r._fulfilTypeList
}

// SetServiceSkuCodeList is ServiceSkuCodeList Setter
// 服务sku，具体的sku列表可以从服务商工作台的类目树获取
func (r *AlibabaSscSupplyplatformServiceabilitySaveAPIRequest) SetServiceSkuCodeList(_serviceSkuCodeList []string) error {
	r._serviceSkuCodeList = _serviceSkuCodeList
	r.Set("service_sku_code_list", _serviceSkuCodeList)
	return nil
}

// GetServiceSkuCodeList ServiceSkuCodeList Getter
func (r AlibabaSscSupplyplatformServiceabilitySaveAPIRequest) GetServiceSkuCodeList() []string {
	return r._serviceSkuCodeList
}

// SetAreaCodeList is AreaCodeList Setter
// 菜鸟地址编码，各级地址均可（全国、省、市、区、街道），根据实际支持的地区填写。当支持的履约类型包含上门时，必填
func (r *AlibabaSscSupplyplatformServiceabilitySaveAPIRequest) SetAreaCodeList(_areaCodeList []int64) error {
	r._areaCodeList = _areaCodeList
	r.Set("area_code_list", _areaCodeList)
	return nil
}

// GetAreaCodeList AreaCodeList Getter
func (r AlibabaSscSupplyplatformServiceabilitySaveAPIRequest) GetAreaCodeList() []int64 {
	return r._areaCodeList
}

// SetProviderType is ProviderType Setter
// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
func (r *AlibabaSscSupplyplatformServiceabilitySaveAPIRequest) SetProviderType(_providerType string) error {
	r._providerType = _providerType
	r.Set("provider_type", _providerType)
	return nil
}

// GetProviderType ProviderType Getter
func (r AlibabaSscSupplyplatformServiceabilitySaveAPIRequest) GetProviderType() string {
	return r._providerType
}

// SetProviderCode is ProviderCode Setter
// 服务提供者code。外部与天猫对接时，自己定义的code用这个字段
func (r *AlibabaSscSupplyplatformServiceabilitySaveAPIRequest) SetProviderCode(_providerCode string) error {
	r._providerCode = _providerCode
	r.Set("provider_code", _providerCode)
	return nil
}

// GetProviderCode ProviderCode Getter
func (r AlibabaSscSupplyplatformServiceabilitySaveAPIRequest) GetProviderCode() string {
	return r._providerCode
}

// SetProviderId is ProviderId Setter
// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
func (r *AlibabaSscSupplyplatformServiceabilitySaveAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r AlibabaSscSupplyplatformServiceabilitySaveAPIRequest) GetProviderId() int64 {
	return r._providerId
}
