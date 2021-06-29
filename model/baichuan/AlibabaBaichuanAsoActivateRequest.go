package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备安装活动激活 API请求
alibaba.baichuan.aso.activate

设备安装活动激活
*/
type AlibabaBaichuanAsoActivateRequest struct {
    model.Params
    // 来源
    source   string
    // 1-tmail,2-taobao
    appId   string
    // 1-android,2-ios
    appOs   int64
    // 设备信息，ios为idfa ，android 为imei + imsi
    deviceInfo   *AsoDeviceInfoDo
}

// 初始化AlibabaBaichuanAsoActivateRequest对象
func NewAlibabaBaichuanAsoActivateRequest() *AlibabaBaichuanAsoActivateRequest{
    return &AlibabaBaichuanAsoActivateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanAsoActivateRequest) GetApiMethodName() string {
    return "alibaba.baichuan.aso.activate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanAsoActivateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Source Setter
// 来源
func (r *AlibabaBaichuanAsoActivateRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r AlibabaBaichuanAsoActivateRequest) GetSource() string {
    return r.source
}
// AppId Setter
// 1-tmail,2-taobao
func (r *AlibabaBaichuanAsoActivateRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

// AppId Getter
func (r AlibabaBaichuanAsoActivateRequest) GetAppId() string {
    return r.appId
}
// AppOs Setter
// 1-android,2-ios
func (r *AlibabaBaichuanAsoActivateRequest) SetAppOs(appOs int64) error {
    r.appOs = appOs
    r.Set("app_os", appOs)
    return nil
}

// AppOs Getter
func (r AlibabaBaichuanAsoActivateRequest) GetAppOs() int64 {
    return r.appOs
}
// DeviceInfo Setter
// 设备信息，ios为idfa ，android 为imei + imsi
func (r *AlibabaBaichuanAsoActivateRequest) SetDeviceInfo(deviceInfo *AsoDeviceInfoDo) error {
    r.deviceInfo = deviceInfo
    r.Set("device_info", deviceInfo)
    return nil
}

// DeviceInfo Getter
func (r AlibabaBaichuanAsoActivateRequest) GetDeviceInfo() *AsoDeviceInfoDo {
    return r.deviceInfo
}
