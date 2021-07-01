package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机设备信息获取 API请求
alibaba.retail.device.info.get

贩卖机设备信息获取
*/
type AlibabaRetailDeviceInfoGetAPIRequest struct {
    model.Params
    // 外部设备ID
    _deviceUuid   string
    // 阿里设备编码
    _deviceCode   string
    // 阿里设备物理ID（32位）
    _deviceSn   string
}

// 初始化AlibabaRetailDeviceInfoGetAPIRequest对象
func NewAlibabaRetailDeviceInfoGetRequest() *AlibabaRetailDeviceInfoGetAPIRequest{
    return &AlibabaRetailDeviceInfoGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailDeviceInfoGetAPIRequest) GetApiMethodName() string {
    return "alibaba.retail.device.info.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailDeviceInfoGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceUuid Setter
// 外部设备ID
func (r *AlibabaRetailDeviceInfoGetAPIRequest) SetDeviceUuid(_deviceUuid string) error {
    r._deviceUuid = _deviceUuid
    r.Set("device_uuid", _deviceUuid)
    return nil
}

// DeviceUuid Getter
func (r AlibabaRetailDeviceInfoGetAPIRequest) GetDeviceUuid() string {
    return r._deviceUuid
}
// DeviceCode Setter
// 阿里设备编码
func (r *AlibabaRetailDeviceInfoGetAPIRequest) SetDeviceCode(_deviceCode string) error {
    r._deviceCode = _deviceCode
    r.Set("device_code", _deviceCode)
    return nil
}

// DeviceCode Getter
func (r AlibabaRetailDeviceInfoGetAPIRequest) GetDeviceCode() string {
    return r._deviceCode
}
// DeviceSn Setter
// 阿里设备物理ID（32位）
func (r *AlibabaRetailDeviceInfoGetAPIRequest) SetDeviceSn(_deviceSn string) error {
    r._deviceSn = _deviceSn
    r.Set("device_sn", _deviceSn)
    return nil
}

// DeviceSn Getter
func (r AlibabaRetailDeviceInfoGetAPIRequest) GetDeviceSn() string {
    return r._deviceSn
}
