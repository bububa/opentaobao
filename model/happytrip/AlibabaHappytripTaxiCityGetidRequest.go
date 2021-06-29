package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取城市id APIRequest
alibaba.happytrip.taxi.city.getid

通过经纬度坐标返回城市id
*/
type AlibabaHappytripTaxiCityGetidRequest struct {
    model.Params

    // 纬度
    lat   string 

    // 经度
    lng   string 

    // 地图类型:amap：高德，默认高德地图
    mapType   string 

}

func NewAlibabaHappytripTaxiCityGetidRequest() *AlibabaHappytripTaxiCityGetidRequest{
    return &AlibabaHappytripTaxiCityGetidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHappytripTaxiCityGetidRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.city.getid"
}

func (r AlibabaHappytripTaxiCityGetidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHappytripTaxiCityGetidRequest) SetLat(lat string) error {
    r.lat = lat
    r.Set("lat", lat)
    return nil
}

func (r AlibabaHappytripTaxiCityGetidRequest) GetLat() string {
    return r.lat
}

func (r *AlibabaHappytripTaxiCityGetidRequest) SetLng(lng string) error {
    r.lng = lng
    r.Set("lng", lng)
    return nil
}

func (r AlibabaHappytripTaxiCityGetidRequest) GetLng() string {
    return r.lng
}

func (r *AlibabaHappytripTaxiCityGetidRequest) SetMapType(mapType string) error {
    r.mapType = mapType
    r.Set("map_type", mapType)
    return nil
}

func (r AlibabaHappytripTaxiCityGetidRequest) GetMapType() string {
    return r.mapType
}

