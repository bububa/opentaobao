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
type AlibabaSscSupplyplatformServicecapacitySaveAPIRequest struct {
    model.Params
    // 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
    _providerType   string
    // 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
    _providerId   int64
    // 目前支持两种。daily 每天容量相同；customize 定制容量，每天都不同
    _mode   string
    // 目前支持两种。day 表示支持的时间粒度为天；hour 时间粒度为小时
    _timeInterval   string
    // 容量数据。根据mode和time interval有不同的格式。具体格式见业务对接文档。
    _capacityData   string
}

// 初始化AlibabaSscSupplyplatformServicecapacitySaveAPIRequest对象
func NewAlibabaSscSupplyplatformServicecapacitySaveRequest() *AlibabaSscSupplyplatformServicecapacitySaveAPIRequest{
    return &AlibabaSscSupplyplatformServicecapacitySaveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServicecapacitySaveAPIRequest) GetApiMethodName() string {
    return "alibaba.ssc.supplyplatform.servicecapacity.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServicecapacitySaveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProviderType Setter
// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
func (r *AlibabaSscSupplyplatformServicecapacitySaveAPIRequest) SetProviderType(_providerType string) error {
    r._providerType = _providerType
    r.Set("provider_type", _providerType)
    return nil
}

// ProviderType Getter
func (r AlibabaSscSupplyplatformServicecapacitySaveAPIRequest) GetProviderType() string {
    return r._providerType
}
// ProviderId Setter
// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
func (r *AlibabaSscSupplyplatformServicecapacitySaveAPIRequest) SetProviderId(_providerId int64) error {
    r._providerId = _providerId
    r.Set("provider_id", _providerId)
    return nil
}

// ProviderId Getter
func (r AlibabaSscSupplyplatformServicecapacitySaveAPIRequest) GetProviderId() int64 {
    return r._providerId
}
// Mode Setter
// 目前支持两种。daily 每天容量相同；customize 定制容量，每天都不同
func (r *AlibabaSscSupplyplatformServicecapacitySaveAPIRequest) SetMode(_mode string) error {
    r._mode = _mode
    r.Set("mode", _mode)
    return nil
}

// Mode Getter
func (r AlibabaSscSupplyplatformServicecapacitySaveAPIRequest) GetMode() string {
    return r._mode
}
// TimeInterval Setter
// 目前支持两种。day 表示支持的时间粒度为天；hour 时间粒度为小时
func (r *AlibabaSscSupplyplatformServicecapacitySaveAPIRequest) SetTimeInterval(_timeInterval string) error {
    r._timeInterval = _timeInterval
    r.Set("time_interval", _timeInterval)
    return nil
}

// TimeInterval Getter
func (r AlibabaSscSupplyplatformServicecapacitySaveAPIRequest) GetTimeInterval() string {
    return r._timeInterval
}
// CapacityData Setter
// 容量数据。根据mode和time interval有不同的格式。具体格式见业务对接文档。
func (r *AlibabaSscSupplyplatformServicecapacitySaveAPIRequest) SetCapacityData(_capacityData string) error {
    r._capacityData = _capacityData
    r.Set("capacity_data", _capacityData)
    return nil
}

// CapacityData Getter
func (r AlibabaSscSupplyplatformServicecapacitySaveAPIRequest) GetCapacityData() string {
    return r._capacityData
}
