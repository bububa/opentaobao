package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备统计数据 APIRequest
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

func NewYunosTvpubadminDeviceStatsRequest() *YunosTvpubadminDeviceStatsRequest{
    return &YunosTvpubadminDeviceStatsRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceStatsRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.stats"
}

func (r YunosTvpubadminDeviceStatsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDeviceStatsRequest) SetFactoryName(factoryName string) error {
    r.factoryName = factoryName
    r.Set("factory_name", factoryName)
    return nil
}

func (r YunosTvpubadminDeviceStatsRequest) GetFactoryName() string {
    return r.factoryName
}

func (r *YunosTvpubadminDeviceStatsRequest) SetDeviceModel(deviceModel string) error {
    r.deviceModel = deviceModel
    r.Set("device_model", deviceModel)
    return nil
}

func (r YunosTvpubadminDeviceStatsRequest) GetDeviceModel() string {
    return r.deviceModel
}

