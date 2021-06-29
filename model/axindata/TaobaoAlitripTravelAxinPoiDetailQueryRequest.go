package axindata

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
景点poi详情查询-阿信 APIRequest
taobao.alitrip.travel.axin.poi.detail.query

景点poi详情查询-阿信
*/
type TaobaoAlitripTravelAxinPoiDetailQueryRequest struct {
    model.Params

    // poiId
    poiId   int64 

}

func NewTaobaoAlitripTravelAxinPoiDetailQueryRequest() *TaobaoAlitripTravelAxinPoiDetailQueryRequest{
    return &TaobaoAlitripTravelAxinPoiDetailQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTravelAxinPoiDetailQueryRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.axin.poi.detail.query"
}

func (r TaobaoAlitripTravelAxinPoiDetailQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTravelAxinPoiDetailQueryRequest) SetPoiId(poiId int64) error {
    r.poiId = poiId
    r.Set("poi_id", poiId)
    return nil
}

func (r TaobaoAlitripTravelAxinPoiDetailQueryRequest) GetPoiId() int64 {
    return r.poiId
}

