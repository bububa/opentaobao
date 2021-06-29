package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务容量删除 API请求
alibaba.ssc.supplyplatform.servicecapacity.delete

服务容量删除
*/
type AlibabaSscSupplyplatformServicecapacityDeleteRequest struct {
    model.Params
    // 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
    providerType   string
    // 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
    providerId   int64
}

// 初始化AlibabaSscSupplyplatformServicecapacityDeleteRequest对象
func NewAlibabaSscSupplyplatformServicecapacityDeleteRequest() *AlibabaSscSupplyplatformServicecapacityDeleteRequest{
    return &AlibabaSscSupplyplatformServicecapacityDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServicecapacityDeleteRequest) GetApiMethodName() string {
    return "alibaba.ssc.supplyplatform.servicecapacity.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServicecapacityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProviderType Setter
// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
func (r *AlibabaSscSupplyplatformServicecapacityDeleteRequest) SetProviderType(providerType string) error {
    r.providerType = providerType
    r.Set("provider_type", providerType)
    return nil
}

// ProviderType Getter
func (r AlibabaSscSupplyplatformServicecapacityDeleteRequest) GetProviderType() string {
    return r.providerType
}
// ProviderId Setter
// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
func (r *AlibabaSscSupplyplatformServicecapacityDeleteRequest) SetProviderId(providerId int64) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

// ProviderId Getter
func (r AlibabaSscSupplyplatformServicecapacityDeleteRequest) GetProviderId() int64 {
    return r.providerId
}
