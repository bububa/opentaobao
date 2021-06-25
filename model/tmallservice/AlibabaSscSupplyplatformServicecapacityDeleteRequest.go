package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务容量删除 APIRequest
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

func NewAlibabaSscSupplyplatformServicecapacityDeleteRequest() *AlibabaSscSupplyplatformServicecapacityDeleteRequest{
    return &AlibabaSscSupplyplatformServicecapacityDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSscSupplyplatformServicecapacityDeleteRequest) GetApiMethodName() string {
    return "alibaba.ssc.supplyplatform.servicecapacity.delete"
}

func (r AlibabaSscSupplyplatformServicecapacityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSscSupplyplatformServicecapacityDeleteRequest) SetProviderType(providerType string) error {
    r.providerType = providerType
    r.Set("provider_type", providerType)
    return nil
}

func (r AlibabaSscSupplyplatformServicecapacityDeleteRequest) GetProviderType() string {
    return r.providerType
}

func (r *AlibabaSscSupplyplatformServicecapacityDeleteRequest) SetProviderId(providerId int64) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

func (r AlibabaSscSupplyplatformServicecapacityDeleteRequest) GetProviderId() int64 {
    return r.providerId
}

