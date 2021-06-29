package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除服务能力 API请求
alibaba.ssc.supplyplatform.serviceability.delete

删除服务能力
*/
type AlibabaSscSupplyplatformServiceabilityDeleteRequest struct {
    model.Params
    // 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
    providerType   string
    // 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
    providerId   int64
}

// 初始化AlibabaSscSupplyplatformServiceabilityDeleteRequest对象
func NewAlibabaSscSupplyplatformServiceabilityDeleteRequest() *AlibabaSscSupplyplatformServiceabilityDeleteRequest{
    return &AlibabaSscSupplyplatformServiceabilityDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServiceabilityDeleteRequest) GetApiMethodName() string {
    return "alibaba.ssc.supplyplatform.serviceability.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServiceabilityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProviderType Setter
// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
func (r *AlibabaSscSupplyplatformServiceabilityDeleteRequest) SetProviderType(providerType string) error {
    r.providerType = providerType
    r.Set("provider_type", providerType)
    return nil
}

// ProviderType Getter
func (r AlibabaSscSupplyplatformServiceabilityDeleteRequest) GetProviderType() string {
    return r.providerType
}
// ProviderId Setter
// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
func (r *AlibabaSscSupplyplatformServiceabilityDeleteRequest) SetProviderId(providerId int64) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

// ProviderId Getter
func (r AlibabaSscSupplyplatformServiceabilityDeleteRequest) GetProviderId() int64 {
    return r.providerId
}
