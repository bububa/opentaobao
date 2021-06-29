package hotelhstdf

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家自主导出相似度高的标准酒店 APIRequest
alitrip.hotel.hstdf.shotel.exportshotel

商家通过给出自己的卖家酒店信息，通过接口可以返回相似度高的标准酒店信息
*/
type AlitripHotelHstdfShotelExportshotelRequest struct {
    model.Params

    // HID，卖家酒店上传到平台后的ID
    hid   int64 

    // 酒店名称，必填
    name   string 

    // 酒店所在行政区划，对应平台ID，为空时会使用经纬度来定位
    cityCode   int64 

    // 电话
    telNumber   string 

    // 经度
    longitude   string 

    // 纬度
    latitude   string 

    // 地址
    address   string 

}

func NewAlitripHotelHstdfShotelExportshotelRequest() *AlitripHotelHstdfShotelExportshotelRequest{
    return &AlitripHotelHstdfShotelExportshotelRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripHotelHstdfShotelExportshotelRequest) GetApiMethodName() string {
    return "alitrip.hotel.hstdf.shotel.exportshotel"
}

func (r AlitripHotelHstdfShotelExportshotelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripHotelHstdfShotelExportshotelRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

func (r AlitripHotelHstdfShotelExportshotelRequest) GetHid() int64 {
    return r.hid
}

func (r *AlitripHotelHstdfShotelExportshotelRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r AlitripHotelHstdfShotelExportshotelRequest) GetName() string {
    return r.name
}

func (r *AlitripHotelHstdfShotelExportshotelRequest) SetCityCode(cityCode int64) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

func (r AlitripHotelHstdfShotelExportshotelRequest) GetCityCode() int64 {
    return r.cityCode
}

func (r *AlitripHotelHstdfShotelExportshotelRequest) SetTelNumber(telNumber string) error {
    r.telNumber = telNumber
    r.Set("tel_number", telNumber)
    return nil
}

func (r AlitripHotelHstdfShotelExportshotelRequest) GetTelNumber() string {
    return r.telNumber
}

func (r *AlitripHotelHstdfShotelExportshotelRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

func (r AlitripHotelHstdfShotelExportshotelRequest) GetLongitude() string {
    return r.longitude
}

func (r *AlitripHotelHstdfShotelExportshotelRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

func (r AlitripHotelHstdfShotelExportshotelRequest) GetLatitude() string {
    return r.latitude
}

func (r *AlitripHotelHstdfShotelExportshotelRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

func (r AlitripHotelHstdfShotelExportshotelRequest) GetAddress() string {
    return r.address
}

