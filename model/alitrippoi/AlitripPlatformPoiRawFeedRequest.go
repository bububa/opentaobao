package alitrippoi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
存储poi原始数据 APIRequest
alitrip.platform.poi.raw.feed

对接外部数据源，外部数据推送poi数据到飞猪
*/
type AlitripPlatformPoiRawFeedRequest struct {
    model.Params

    // poi存储参数
    param0   *TripPoiRawSaveParam 

}

func NewAlitripPlatformPoiRawFeedRequest() *AlitripPlatformPoiRawFeedRequest{
    return &AlitripPlatformPoiRawFeedRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripPlatformPoiRawFeedRequest) GetApiMethodName() string {
    return "alitrip.platform.poi.raw.feed"
}

func (r AlitripPlatformPoiRawFeedRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripPlatformPoiRawFeedRequest) SetParam0(param0 *TripPoiRawSaveParam) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlitripPlatformPoiRawFeedRequest) GetParam0() *TripPoiRawSaveParam {
    return r.param0
}

