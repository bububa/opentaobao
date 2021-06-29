package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务库存查询 API请求
alibaba.ssc.supplyplatform.service.inventory.query

查询服务库存。需要保存服务容量成功后，才能进行查询，参数中的provider信息（provider_id和provider_type）与alibaba.ssc.supplyplatform.servicecapacity.save接口中保持一致。
*/
type AlibabaSscSupplyplatformServiceInventoryQueryRequest struct {
    model.Params
    // 查询开始日期。yyyy-MM-dd格式
    startDay   string
    // 查询结束日期。与start间隔不能超过31天。yyyy-MM-dd格式
    endDay   string
    // 服务提供者类型。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
    providerType   string
    // 服务提供者类型。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
    providerId   int64
}

// 初始化AlibabaSscSupplyplatformServiceInventoryQueryRequest对象
func NewAlibabaSscSupplyplatformServiceInventoryQueryRequest() *AlibabaSscSupplyplatformServiceInventoryQueryRequest{
    return &AlibabaSscSupplyplatformServiceInventoryQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServiceInventoryQueryRequest) GetApiMethodName() string {
    return "alibaba.ssc.supplyplatform.service.inventory.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServiceInventoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartDay Setter
// 查询开始日期。yyyy-MM-dd格式
func (r *AlibabaSscSupplyplatformServiceInventoryQueryRequest) SetStartDay(startDay string) error {
    r.startDay = startDay
    r.Set("start_day", startDay)
    return nil
}

// StartDay Getter
func (r AlibabaSscSupplyplatformServiceInventoryQueryRequest) GetStartDay() string {
    return r.startDay
}
// EndDay Setter
// 查询结束日期。与start间隔不能超过31天。yyyy-MM-dd格式
func (r *AlibabaSscSupplyplatformServiceInventoryQueryRequest) SetEndDay(endDay string) error {
    r.endDay = endDay
    r.Set("end_day", endDay)
    return nil
}

// EndDay Getter
func (r AlibabaSscSupplyplatformServiceInventoryQueryRequest) GetEndDay() string {
    return r.endDay
}
// ProviderType Setter
// 服务提供者类型。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
func (r *AlibabaSscSupplyplatformServiceInventoryQueryRequest) SetProviderType(providerType string) error {
    r.providerType = providerType
    r.Set("provider_type", providerType)
    return nil
}

// ProviderType Getter
func (r AlibabaSscSupplyplatformServiceInventoryQueryRequest) GetProviderType() string {
    return r.providerType
}
// ProviderId Setter
// 服务提供者类型。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
func (r *AlibabaSscSupplyplatformServiceInventoryQueryRequest) SetProviderId(providerId int64) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

// ProviderId Getter
func (r AlibabaSscSupplyplatformServiceInventoryQueryRequest) GetProviderId() int64 {
    return r.providerId
}
