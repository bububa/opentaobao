package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询app在设备上的安装信息 APIRequest
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

func NewAlibabaBaichuanAsoQueryRequest() *AlibabaBaichuanAsoQueryRequest{
    return &AlibabaBaichuanAsoQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaBaichuanAsoQueryRequest) GetApiMethodName() string {
    return "alibaba.baichuan.aso.query"
}

func (r AlibabaBaichuanAsoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaBaichuanAsoQueryRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

func (r AlibabaBaichuanAsoQueryRequest) GetAppId() string {
    return r.appId
}

func (r *AlibabaBaichuanAsoQueryRequest) SetAppOs(appOs int64) error {
    r.appOs = appOs
    r.Set("app_os", appOs)
    return nil
}

func (r AlibabaBaichuanAsoQueryRequest) GetAppOs() int64 {
    return r.appOs
}

func (r *AlibabaBaichuanAsoQueryRequest) SetDeviceInfoList(deviceInfoList []ASODeviceInfoDO) error {
    r.deviceInfoList = deviceInfoList
    r.Set("device_info_list", deviceInfoList)
    return nil
}

func (r AlibabaBaichuanAsoQueryRequest) GetDeviceInfoList() []ASODeviceInfoDO {
    return r.deviceInfoList
}

