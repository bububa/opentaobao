package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取城市id API请求
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

// 初始化AlibabaHappytripTaxiCityGetidRequest对象
func NewAlibabaHappytripTaxiCityGetidRequest() *AlibabaHappytripTaxiCityGetidRequest{
    return &AlibabaHappytripTaxiCityGetidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiCityGetidRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.city.getid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiCityGetidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Lat Setter
// 纬度
func (r *AlibabaHappytripTaxiCityGetidRequest) SetLat(lat string) error {
    r.lat = lat
    r.Set("lat", lat)
    return nil
}

// Lat Getter
func (r AlibabaHappytripTaxiCityGetidRequest) GetLat() string {
    return r.lat
}
// Lng Setter
// 经度
func (r *AlibabaHappytripTaxiCityGetidRequest) SetLng(lng string) error {
    r.lng = lng
    r.Set("lng", lng)
    return nil
}

// Lng Getter
func (r AlibabaHappytripTaxiCityGetidRequest) GetLng() string {
    return r.lng
}
// MapType Setter
// 地图类型:amap：高德，默认高德地图
func (r *AlibabaHappytripTaxiCityGetidRequest) SetMapType(mapType string) error {
    r.mapType = mapType
    r.Set("map_type", mapType)
    return nil
}

// MapType Getter
func (r AlibabaHappytripTaxiCityGetidRequest) GetMapType() string {
    return r.mapType
}
