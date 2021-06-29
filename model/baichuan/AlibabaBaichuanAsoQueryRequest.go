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
    appId   string
    // 1-android,2-ios
    appOs   int64
    // 设备信息，ios为idfa ，android 为imei + imsi
    deviceInfoList   []ASODeviceInfoDO
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
func (r *AlibabaBaichuanAsoQueryRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

// AppId Getter
func (r AlibabaBaichuanAsoQueryRequest) GetAppId() string {
    return r.appId
}
// AppOs Setter
// 1-android,2-ios
func (r *AlibabaBaichuanAsoQueryRequest) SetAppOs(appOs int64) error {
    r.appOs = appOs
    r.Set("app_os", appOs)
    return nil
}

// AppOs Getter
func (r AlibabaBaichuanAsoQueryRequest) GetAppOs() int64 {
    return r.appOs
}
// DeviceInfoList Setter
// 设备信息，ios为idfa ，android 为imei + imsi
func (r *AlibabaBaichuanAsoQueryRequest) SetDeviceInfoList(deviceInfoList []ASODeviceInfoDO) error {
    r.deviceInfoList = deviceInfoList
    r.Set("device_info_list", deviceInfoList)
    return nil
}

// DeviceInfoList Getter
func (r AlibabaBaichuanAsoQueryRequest) GetDeviceInfoList() []ASODeviceInfoDO {
    return r.deviceInfoList
}
