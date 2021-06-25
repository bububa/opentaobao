package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除服务能力 APIRequest
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

func NewAlibabaSscSupplyplatformServiceabilityDeleteRequest() *AlibabaSscSupplyplatformServiceabilityDeleteRequest{
    return &AlibabaSscSupplyplatformServiceabilityDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSscSupplyplatformServiceabilityDeleteRequest) GetApiMethodName() string {
    return "alibaba.ssc.supplyplatform.serviceability.delete"
}

func (r AlibabaSscSupplyplatformServiceabilityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSscSupplyplatformServiceabilityDeleteRequest) SetProviderType(providerType string) error {
    r.providerType = providerType
    r.Set("provider_type", providerType)
    return nil
}

func (r AlibabaSscSupplyplatformServiceabilityDeleteRequest) GetProviderType() string {
    return r.providerType
}

func (r *AlibabaSscSupplyplatformServiceabilityDeleteRequest) SetProviderId(providerId int64) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

func (r AlibabaSscSupplyplatformServiceabilityDeleteRequest) GetProviderId() int64 {
    return r.providerId
}

