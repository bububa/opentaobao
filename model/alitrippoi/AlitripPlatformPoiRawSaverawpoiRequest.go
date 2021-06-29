package alitrippoi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
POI开放存储能力 APIRequest
alitrip.platform.poi.raw.saverawpoi

POI开放存储提供离线/在线/纬错更新的能力
*/
type AlitripPlatformPoiRawSaverawpoiRequest struct {
    model.Params

    // poi存储参数
    tripPoiRawSaveParam   *TripPoiRawSaveParamV2 

}

func NewAlitripPlatformPoiRawSaverawpoiRequest() *AlitripPlatformPoiRawSaverawpoiRequest{
    return &AlitripPlatformPoiRawSaverawpoiRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripPlatformPoiRawSaverawpoiRequest) GetApiMethodName() string {
    return "alitrip.platform.poi.raw.saverawpoi"
}

func (r AlitripPlatformPoiRawSaverawpoiRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripPlatformPoiRawSaverawpoiRequest) SetTripPoiRawSaveParam(tripPoiRawSaveParam *TripPoiRawSaveParamV2) error {
    r.tripPoiRawSaveParam = tripPoiRawSaveParam
    r.Set("trip_poi_raw_save_param", tripPoiRawSaveParam)
    return nil
}

func (r AlitripPlatformPoiRawSaverawpoiRequest) GetTripPoiRawSaveParam() *TripPoiRawSaveParamV2 {
    return r.tripPoiRawSaveParam
}

