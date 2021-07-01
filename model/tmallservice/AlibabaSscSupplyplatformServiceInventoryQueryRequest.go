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
type AlibabaSscSupplyplatformServiceInventoryQueryAPIRequest struct {
    model.Params
    // 查询开始日期。yyyy-MM-dd格式
    _startDay   string
    // 查询结束日期。与start间隔不能超过31天。yyyy-MM-dd格式
    _endDay   string
    // 服务提供者类型。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
    _providerType   string
    // 服务提供者类型。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
    _providerId   int64
}

// 初始化AlibabaSscSupplyplatformServiceInventoryQueryAPIRequest对象
func NewAlibabaSscSupplyplatformServiceInventoryQueryRequest() *AlibabaSscSupplyplatformServiceInventoryQueryAPIRequest{
    return &AlibabaSscSupplyplatformServiceInventoryQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServiceInventoryQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.ssc.supplyplatform.service.inventory.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServiceInventoryQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartDay Setter
// 查询开始日期。yyyy-MM-dd格式
func (r *AlibabaSscSupplyplatformServiceInventoryQueryAPIRequest) SetStartDay(_startDay string) error {
    r._startDay = _startDay
    r.Set("start_day", _startDay)
    return nil
}

// StartDay Getter
func (r AlibabaSscSupplyplatformServiceInventoryQueryAPIRequest) GetStartDay() string {
    return r._startDay
}
// EndDay Setter
// 查询结束日期。与start间隔不能超过31天。yyyy-MM-dd格式
func (r *AlibabaSscSupplyplatformServiceInventoryQueryAPIRequest) SetEndDay(_endDay string) error {
    r._endDay = _endDay
    r.Set("end_day", _endDay)
    return nil
}

// EndDay Getter
func (r AlibabaSscSupplyplatformServiceInventoryQueryAPIRequest) GetEndDay() string {
    return r._endDay
}
// ProviderType Setter
// 服务提供者类型。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
func (r *AlibabaSscSupplyplatformServiceInventoryQueryAPIRequest) SetProviderType(_providerType string) error {
    r._providerType = _providerType
    r.Set("provider_type", _providerType)
    return nil
}

// ProviderType Getter
func (r AlibabaSscSupplyplatformServiceInventoryQueryAPIRequest) GetProviderType() string {
    return r._providerType
}
// ProviderId Setter
// 服务提供者类型。参考alibaba.ssc.supplyplatform.servicecapacity.save入参
func (r *AlibabaSscSupplyplatformServiceInventoryQueryAPIRequest) SetProviderId(_providerId int64) error {
    r._providerId = _providerId
    r.Set("provider_id", _providerId)
    return nil
}

// ProviderId Getter
func (r AlibabaSscSupplyplatformServiceInventoryQueryAPIRequest) GetProviderId() int64 {
    return r._providerId
}
