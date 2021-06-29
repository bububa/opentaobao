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
    _source   string
    // 1-tmail,2-taobao
    _appId   string
    // 1-android,2-ios
    _appOs   int64
    // 设备信息，ios为idfa ，android 为imei + imsi
    _deviceInfo   *AsoDeviceInfoDO
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
func (r *AlibabaBaichuanAsoActivateRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r AlibabaBaichuanAsoActivateRequest) GetSource() string {
    return r._source
}
// AppId Setter
// 1-tmail,2-taobao
func (r *AlibabaBaichuanAsoActivateRequest) SetAppId(_appId string) error {
    r._appId = _appId
    r.Set("app_id", _appId)
    return nil
}

// AppId Getter
func (r AlibabaBaichuanAsoActivateRequest) GetAppId() string {
    return r._appId
}
// AppOs Setter
// 1-android,2-ios
func (r *AlibabaBaichuanAsoActivateRequest) SetAppOs(_appOs int64) error {
    r._appOs = _appOs
    r.Set("app_os", _appOs)
    return nil
}

// AppOs Getter
func (r AlibabaBaichuanAsoActivateRequest) GetAppOs() int64 {
    return r._appOs
}
// DeviceInfo Setter
// 设备信息，ios为idfa ，android 为imei + imsi
func (r *AlibabaBaichuanAsoActivateRequest) SetDeviceInfo(_deviceInfo *AsoDeviceInfoDO) error {
    r._deviceInfo = _deviceInfo
    r.Set("device_info", _deviceInfo)
    return nil
}

// DeviceInfo Getter
func (r AlibabaBaichuanAsoActivateRequest) GetDeviceInfo() *AsoDeviceInfoDO {
    return r._deviceInfo
}
