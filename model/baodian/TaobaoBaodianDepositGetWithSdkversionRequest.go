package baodian

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户宝点信息（带sdk版本，已迁移） APIRequest
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

func NewTaobaoBaodianDepositGetWithSdkversionRequest() *TaobaoBaodianDepositGetWithSdkversionRequest{
    return &TaobaoBaodianDepositGetWithSdkversionRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaodianDepositGetWithSdkversionRequest) GetApiMethodName() string {
    return "taobao.baodian.deposit.get.with.sdkversion"
}

func (r TaobaoBaodianDepositGetWithSdkversionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaodianDepositGetWithSdkversionRequest) SetDeviceModel(deviceModel string) error {
    r.deviceModel = deviceModel
    r.Set("device_model", deviceModel)
    return nil
}

func (r TaobaoBaodianDepositGetWithSdkversionRequest) GetDeviceModel() string {
    return r.deviceModel
}

func (r *TaobaoBaodianDepositGetWithSdkversionRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r TaobaoBaodianDepositGetWithSdkversionRequest) GetUuid() string {
    return r.uuid
}

func (r *TaobaoBaodianDepositGetWithSdkversionRequest) SetSdkVersion(sdkVersion string) error {
    r.sdkVersion = sdkVersion
    r.Set("sdk_version", sdkVersion)
    return nil
}

func (r TaobaoBaodianDepositGetWithSdkversionRequest) GetSdkVersion() string {
    return r.sdkVersion
}

