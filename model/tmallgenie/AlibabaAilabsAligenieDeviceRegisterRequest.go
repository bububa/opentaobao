package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵开放平台获取设备秘钥 API请求
alibaba.ailabs.aligenie.device.register

向天猫精灵inside平台注册设备mac地址，并获取设备的唯一密钥
*/
type AlibabaAilabsAligenieDeviceRegisterRequest struct {
    model.Params
    // 设备id
    _deviceId   int64
    // mac区段脚本
    _macSections   string
}

// 初始化AlibabaAilabsAligenieDeviceRegisterRequest对象
func NewAlibabaAilabsAligenieDeviceRegisterRequest() *AlibabaAilabsAligenieDeviceRegisterRequest{
    return &AlibabaAilabsAligenieDeviceRegisterRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieDeviceRegisterRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.device.register"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieDeviceRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceId Setter
// 设备id
func (r *AlibabaAilabsAligenieDeviceRegisterRequest) SetDeviceId(_deviceId int64) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r AlibabaAilabsAligenieDeviceRegisterRequest) GetDeviceId() int64 {
    return r._deviceId
}
// MacSections Setter
// mac区段脚本
func (r *AlibabaAilabsAligenieDeviceRegisterRequest) SetMacSections(_macSections string) error {
    r._macSections = _macSections
    r.Set("mac_sections", _macSections)
    return nil
}

// MacSections Getter
func (r AlibabaAilabsAligenieDeviceRegisterRequest) GetMacSections() string {
    return r._macSections
}
