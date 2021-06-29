package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备统计数据 API请求
yunos.tvpubadmin.device.stats

获取设备统计数据
*/
type YunosTvpubadminDeviceStatsRequest struct {
    model.Params
    // 厂商名称
    factoryName   string
    // 设备型号
    deviceModel   string
}

// 初始化YunosTvpubadminDeviceStatsRequest对象
func NewYunosTvpubadminDeviceStatsRequest() *YunosTvpubadminDeviceStatsRequest{
    return &YunosTvpubadminDeviceStatsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceStatsRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.stats"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceStatsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FactoryName Setter
// 厂商名称
func (r *YunosTvpubadminDeviceStatsRequest) SetFactoryName(factoryName string) error {
    r.factoryName = factoryName
    r.Set("factory_name", factoryName)
    return nil
}

// FactoryName Getter
func (r YunosTvpubadminDeviceStatsRequest) GetFactoryName() string {
    return r.factoryName
}
// DeviceModel Setter
// 设备型号
func (r *YunosTvpubadminDeviceStatsRequest) SetDeviceModel(deviceModel string) error {
    r.deviceModel = deviceModel
    r.Set("device_model", deviceModel)
    return nil
}

// DeviceModel Getter
func (r YunosTvpubadminDeviceStatsRequest) GetDeviceModel() string {
    return r.deviceModel
}
