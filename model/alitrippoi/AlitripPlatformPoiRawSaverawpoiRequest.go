package alitrippoi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
POI开放存储能力 API请求
alitrip.platform.poi.raw.saverawpoi

POI开放存储提供离线/在线/纬错更新的能力
*/
type AlitripPlatformPoiRawSaverawpoiAPIRequest struct {
    model.Params
    // poi存储参数
    _tripPoiRawSaveParam   *TripPoiRawSaveParamV2
}

// 初始化AlitripPlatformPoiRawSaverawpoiAPIRequest对象
func NewAlitripPlatformPoiRawSaverawpoiRequest() *AlitripPlatformPoiRawSaverawpoiAPIRequest{
    return &AlitripPlatformPoiRawSaverawpoiAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripPlatformPoiRawSaverawpoiAPIRequest) GetApiMethodName() string {
    return "alitrip.platform.poi.raw.saverawpoi"
}

// IRequest interface 方法, 获取API参数
func (r AlitripPlatformPoiRawSaverawpoiAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TripPoiRawSaveParam Setter
// poi存储参数
func (r *AlitripPlatformPoiRawSaverawpoiAPIRequest) SetTripPoiRawSaveParam(_tripPoiRawSaveParam *TripPoiRawSaveParamV2) error {
    r._tripPoiRawSaveParam = _tripPoiRawSaveParam
    r.Set("trip_poi_raw_save_param", _tripPoiRawSaveParam)
    return nil
}

// TripPoiRawSaveParam Getter
func (r AlitripPlatformPoiRawSaverawpoiAPIRequest) GetTripPoiRawSaveParam() *TripPoiRawSaveParamV2 {
    return r._tripPoiRawSaveParam
}
