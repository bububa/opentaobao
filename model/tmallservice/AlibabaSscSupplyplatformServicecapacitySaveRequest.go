package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
保存服务容量 API请求
alibaba.ssc.supplyplatform.servicecapacity.save

保存服务容量
*/
type AlibabaSscSupplyplatformServicecapacitySaveRequest struct {
    model.Params
    // 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
    providerType   string
    // 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
    providerId   int64
    // 目前支持两种。daily 每天容量相同；customize 定制容量，每天都不同
    mode   string
    // 目前支持两种。day 表示支持的时间粒度为天；hour 时间粒度为小时
    timeInterval   string
    // 容量数据。根据mode和time interval有不同的格式。具体格式见业务对接文档。
    capacityData   string
}

// 初始化AlibabaSscSupplyplatformServicecapacitySaveRequest对象
func NewAlibabaSscSupplyplatformServicecapacitySaveRequest() *AlibabaSscSupplyplatformServicecapacitySaveRequest{
    return &AlibabaSscSupplyplatformServicecapacitySaveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServicecapacitySaveRequest) GetApiMethodName() string {
    return "alibaba.ssc.supplyplatform.servicecapacity.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServicecapacitySaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProviderType Setter
// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
func (r *AlibabaSscSupplyplatformServicecapacitySaveRequest) SetProviderType(providerType string) error {
    r.providerType = providerType
    r.Set("provider_type", providerType)
    return nil
}

// ProviderType Getter
func (r AlibabaSscSupplyplatformServicecapacitySaveRequest) GetProviderType() string {
    return r.providerType
}
// ProviderId Setter
// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
func (r *AlibabaSscSupplyplatformServicecapacitySaveRequest) SetProviderId(providerId int64) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

// ProviderId Getter
func (r AlibabaSscSupplyplatformServicecapacitySaveRequest) GetProviderId() int64 {
    return r.providerId
}
// Mode Setter
// 目前支持两种。daily 每天容量相同；customize 定制容量，每天都不同
func (r *AlibabaSscSupplyplatformServicecapacitySaveRequest) SetMode(mode string) error {
    r.mode = mode
    r.Set("mode", mode)
    return nil
}

// Mode Getter
func (r AlibabaSscSupplyplatformServicecapacitySaveRequest) GetMode() string {
    return r.mode
}
// TimeInterval Setter
// 目前支持两种。day 表示支持的时间粒度为天；hour 时间粒度为小时
func (r *AlibabaSscSupplyplatformServicecapacitySaveRequest) SetTimeInterval(timeInterval string) error {
    r.timeInterval = timeInterval
    r.Set("time_interval", timeInterval)
    return nil
}

// TimeInterval Getter
func (r AlibabaSscSupplyplatformServicecapacitySaveRequest) GetTimeInterval() string {
    return r.timeInterval
}
// CapacityData Setter
// 容量数据。根据mode和time interval有不同的格式。具体格式见业务对接文档。
func (r *AlibabaSscSupplyplatformServicecapacitySaveRequest) SetCapacityData(capacityData string) error {
    r.capacityData = capacityData
    r.Set("capacity_data", capacityData)
    return nil
}

// CapacityData Getter
func (r AlibabaSscSupplyplatformServicecapacitySaveRequest) GetCapacityData() string {
    return r.capacityData
}
