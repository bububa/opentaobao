package alitrippoi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
存储poi原始数据 API请求
alitrip.platform.poi.raw.feed

对接外部数据源，外部数据推送poi数据到飞猪
*/
type AlitripPlatformPoiRawFeedRequest struct {
    model.Params
    // poi存储参数
    _param0   *TripPoiRawSaveParam
}

// 初始化AlitripPlatformPoiRawFeedRequest对象
func NewAlitripPlatformPoiRawFeedRequest() *AlitripPlatformPoiRawFeedRequest{
    return &AlitripPlatformPoiRawFeedRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripPlatformPoiRawFeedRequest) GetApiMethodName() string {
    return "alitrip.platform.poi.raw.feed"
}

// IRequest interface 方法, 获取API参数
func (r AlitripPlatformPoiRawFeedRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// poi存储参数
func (r *AlitripPlatformPoiRawFeedRequest) SetParam0(_param0 *TripPoiRawSaveParam) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlitripPlatformPoiRawFeedRequest) GetParam0() *TripPoiRawSaveParam {
    return r._param0
}
