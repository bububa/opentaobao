package cainiaolocker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
事件回传接口 APIRequest
cainiao.endpoint.locker.top.order.tracking.new

用于合作公司通知入站、出站信息到菜鸟，共定义了5个操作，1个投件操作，3个取件操作，1个超时提醒。
*/
type CainiaoEndpointLockerTopOrderTrackingNewRequest struct {
    model.Params

    // 回传信息
    trackInfo   *CollectTrackingInfo 

}

func NewCainiaoEndpointLockerTopOrderTrackingNewRequest() *CainiaoEndpointLockerTopOrderTrackingNewRequest{
    return &CainiaoEndpointLockerTopOrderTrackingNewRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoEndpointLockerTopOrderTrackingNewRequest) GetApiMethodName() string {
    return "cainiao.endpoint.locker.top.order.tracking.new"
}

func (r CainiaoEndpointLockerTopOrderTrackingNewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoEndpointLockerTopOrderTrackingNewRequest) SetTrackInfo(trackInfo *CollectTrackingInfo) error {
    r.trackInfo = trackInfo
    r.Set("track_info", trackInfo)
    return nil
}

func (r CainiaoEndpointLockerTopOrderTrackingNewRequest) GetTrackInfo() *CollectTrackingInfo {
    return r.trackInfo
}

