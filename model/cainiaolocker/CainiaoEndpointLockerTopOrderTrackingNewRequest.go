package cainiaolocker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
事件回传接口 API请求
cainiao.endpoint.locker.top.order.tracking.new

用于合作公司通知入站、出站信息到菜鸟，共定义了5个操作，1个投件操作，3个取件操作，1个超时提醒。
*/
type CainiaoEndpointLockerTopOrderTrackingNewRequest struct {
    model.Params
    // 回传信息
    _trackInfo   *CollectTrackingInfo
}

// 初始化CainiaoEndpointLockerTopOrderTrackingNewRequest对象
func NewCainiaoEndpointLockerTopOrderTrackingNewRequest() *CainiaoEndpointLockerTopOrderTrackingNewRequest{
    return &CainiaoEndpointLockerTopOrderTrackingNewRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoEndpointLockerTopOrderTrackingNewRequest) GetApiMethodName() string {
    return "cainiao.endpoint.locker.top.order.tracking.new"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoEndpointLockerTopOrderTrackingNewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TrackInfo Setter
// 回传信息
func (r *CainiaoEndpointLockerTopOrderTrackingNewRequest) SetTrackInfo(_trackInfo *CollectTrackingInfo) error {
    r._trackInfo = _trackInfo
    r.Set("track_info", _trackInfo)
    return nil
}

// TrackInfo Getter
func (r CainiaoEndpointLockerTopOrderTrackingNewRequest) GetTrackInfo() *CollectTrackingInfo {
    return r._trackInfo
}
