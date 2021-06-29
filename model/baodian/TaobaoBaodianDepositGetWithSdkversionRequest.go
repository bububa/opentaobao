package baodian

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户宝点信息（带sdk版本，已迁移） API请求
taobao.baodian.deposit.get.with.sdkversion

获取用户宝点信息（带sdk版本，已迁移）
*/
type TaobaoBaodianDepositGetWithSdkversionRequest struct {
    model.Params
    // 设备型号
    deviceModel   string
    // uuid
    uuid   string
    // sdk版本
    sdkVersion   string
}

// 初始化TaobaoBaodianDepositGetWithSdkversionRequest对象
func NewTaobaoBaodianDepositGetWithSdkversionRequest() *TaobaoBaodianDepositGetWithSdkversionRequest{
    return &TaobaoBaodianDepositGetWithSdkversionRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaodianDepositGetWithSdkversionRequest) GetApiMethodName() string {
    return "taobao.baodian.deposit.get.with.sdkversion"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaodianDepositGetWithSdkversionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceModel Setter
// 设备型号
func (r *TaobaoBaodianDepositGetWithSdkversionRequest) SetDeviceModel(deviceModel string) error {
    r.deviceModel = deviceModel
    r.Set("device_model", deviceModel)
    return nil
}

// DeviceModel Getter
func (r TaobaoBaodianDepositGetWithSdkversionRequest) GetDeviceModel() string {
    return r.deviceModel
}
// Uuid Setter
// uuid
func (r *TaobaoBaodianDepositGetWithSdkversionRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r TaobaoBaodianDepositGetWithSdkversionRequest) GetUuid() string {
    return r.uuid
}
// SdkVersion Setter
// sdk版本
func (r *TaobaoBaodianDepositGetWithSdkversionRequest) SetSdkVersion(sdkVersion string) error {
    r.sdkVersion = sdkVersion
    r.Set("sdk_version", sdkVersion)
    return nil
}

// SdkVersion Getter
func (r TaobaoBaodianDepositGetWithSdkversionRequest) GetSdkVersion() string {
    return r.sdkVersion
}
