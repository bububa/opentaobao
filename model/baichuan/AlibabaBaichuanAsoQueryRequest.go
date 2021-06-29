package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询app在设备上的安装信息 API请求
alibaba.baichuan.aso.query

查询app在设备上的安装信息
*/
type AlibabaBaichuanAsoQueryRequest struct {
    model.Params
    // 1-tmail,2-taobao
    _appId   string
    // 1-android,2-ios
    _appOs   int64
    // 设备信息，ios为idfa ，android 为imei + imsi
    _deviceInfoList   []ASODeviceInfoDO
}

// 初始化AlibabaBaichuanAsoQueryRequest对象
func NewAlibabaBaichuanAsoQueryRequest() *AlibabaBaichuanAsoQueryRequest{
    return &AlibabaBaichuanAsoQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanAsoQueryRequest) GetApiMethodName() string {
    return "alibaba.baichuan.aso.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanAsoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppId Setter
// 1-tmail,2-taobao
func (r *AlibabaBaichuanAsoQueryRequest) SetAppId(_appId string) error {
    r._appId = _appId
    r.Set("app_id", _appId)
    return nil
}

// AppId Getter
func (r AlibabaBaichuanAsoQueryRequest) GetAppId() string {
    return r._appId
}
// AppOs Setter
// 1-android,2-ios
func (r *AlibabaBaichuanAsoQueryRequest) SetAppOs(_appOs int64) error {
    r._appOs = _appOs
    r.Set("app_os", _appOs)
    return nil
}

// AppOs Getter
func (r AlibabaBaichuanAsoQueryRequest) GetAppOs() int64 {
    return r._appOs
}
// DeviceInfoList Setter
// 设备信息，ios为idfa ，android 为imei + imsi
func (r *AlibabaBaichuanAsoQueryRequest) SetDeviceInfoList(_deviceInfoList []ASODeviceInfoDO) error {
    r._deviceInfoList = _deviceInfoList
    r.Set("device_info_list", _deviceInfoList)
    return nil
}

// DeviceInfoList Getter
func (r AlibabaBaichuanAsoQueryRequest) GetDeviceInfoList() []ASODeviceInfoDO {
    return r._deviceInfoList
}
