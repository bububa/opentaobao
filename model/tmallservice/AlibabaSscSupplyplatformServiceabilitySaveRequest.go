package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
保存服务能力 API请求
alibaba.ssc.supplyplatform.serviceability.save

保存服务能力
*/
type AlibabaSscSupplyplatformServiceabilitySaveRequest struct {
    model.Params
    // 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
    _providerType   string
    // 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
    _providerId   int64
    // 目前包含三种。in_store 到店；at_home 上门；transmit_service 寄修。请根据实际支持的履约类型填写
    _fulfilTypeList   []string
    // 服务sku，具体的sku列表可以从服务商工作台的类目树获取
    _serviceSkuCodeList   []string
    // 菜鸟地址编码，各级地址均可（全国、省、市、区、街道），根据实际支持的地区填写。当支持的履约类型包含上门时，必填
    _areaCodeList   []int64
}

// 初始化AlibabaSscSupplyplatformServiceabilitySaveRequest对象
func NewAlibabaSscSupplyplatformServiceabilitySaveRequest() *AlibabaSscSupplyplatformServiceabilitySaveRequest{
    return &AlibabaSscSupplyplatformServiceabilitySaveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServiceabilitySaveRequest) GetApiMethodName() string {
    return "alibaba.ssc.supplyplatform.serviceability.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServiceabilitySaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProviderType Setter
// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
func (r *AlibabaSscSupplyplatformServiceabilitySaveRequest) SetProviderType(_providerType string) error {
    r._providerType = _providerType
    r.Set("provider_type", _providerType)
    return nil
}

// ProviderType Getter
func (r AlibabaSscSupplyplatformServiceabilitySaveRequest) GetProviderType() string {
    return r._providerType
}
// ProviderId Setter
// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
func (r *AlibabaSscSupplyplatformServiceabilitySaveRequest) SetProviderId(_providerId int64) error {
    r._providerId = _providerId
    r.Set("provider_id", _providerId)
    return nil
}

// ProviderId Getter
func (r AlibabaSscSupplyplatformServiceabilitySaveRequest) GetProviderId() int64 {
    return r._providerId
}
// FulfilTypeList Setter
// 目前包含三种。in_store 到店；at_home 上门；transmit_service 寄修。请根据实际支持的履约类型填写
func (r *AlibabaSscSupplyplatformServiceabilitySaveRequest) SetFulfilTypeList(_fulfilTypeList []string) error {
    r._fulfilTypeList = _fulfilTypeList
    r.Set("fulfil_type_list", _fulfilTypeList)
    return nil
}

// FulfilTypeList Getter
func (r AlibabaSscSupplyplatformServiceabilitySaveRequest) GetFulfilTypeList() []string {
    return r._fulfilTypeList
}
// ServiceSkuCodeList Setter
// 服务sku，具体的sku列表可以从服务商工作台的类目树获取
func (r *AlibabaSscSupplyplatformServiceabilitySaveRequest) SetServiceSkuCodeList(_serviceSkuCodeList []string) error {
    r._serviceSkuCodeList = _serviceSkuCodeList
    r.Set("service_sku_code_list", _serviceSkuCodeList)
    return nil
}

// ServiceSkuCodeList Getter
func (r AlibabaSscSupplyplatformServiceabilitySaveRequest) GetServiceSkuCodeList() []string {
    return r._serviceSkuCodeList
}
// AreaCodeList Setter
// 菜鸟地址编码，各级地址均可（全国、省、市、区、街道），根据实际支持的地区填写。当支持的履约类型包含上门时，必填
func (r *AlibabaSscSupplyplatformServiceabilitySaveRequest) SetAreaCodeList(_areaCodeList []int64) error {
    r._areaCodeList = _areaCodeList
    r.Set("area_code_list", _areaCodeList)
    return nil
}

// AreaCodeList Getter
func (r AlibabaSscSupplyplatformServiceabilitySaveRequest) GetAreaCodeList() []int64 {
    return r._areaCodeList
}
