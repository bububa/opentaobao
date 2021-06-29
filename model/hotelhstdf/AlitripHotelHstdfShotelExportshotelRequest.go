package hotelhstdf

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家自主导出相似度高的标准酒店 API请求
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

// 初始化AlitripHotelHstdfShotelExportshotelRequest对象
func NewAlitripHotelHstdfShotelExportshotelRequest() *AlitripHotelHstdfShotelExportshotelRequest{
    return &AlitripHotelHstdfShotelExportshotelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfShotelExportshotelRequest) GetApiMethodName() string {
    return "alitrip.hotel.hstdf.shotel.exportshotel"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfShotelExportshotelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Hid Setter
// HID，卖家酒店上传到平台后的ID
func (r *AlitripHotelHstdfShotelExportshotelRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

// Hid Getter
func (r AlitripHotelHstdfShotelExportshotelRequest) GetHid() int64 {
    return r.hid
}
// Name Setter
// 酒店名称，必填
func (r *AlitripHotelHstdfShotelExportshotelRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r AlitripHotelHstdfShotelExportshotelRequest) GetName() string {
    return r.name
}
// CityCode Setter
// 酒店所在行政区划，对应平台ID，为空时会使用经纬度来定位
func (r *AlitripHotelHstdfShotelExportshotelRequest) SetCityCode(cityCode int64) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

// CityCode Getter
func (r AlitripHotelHstdfShotelExportshotelRequest) GetCityCode() int64 {
    return r.cityCode
}
// TelNumber Setter
// 电话
func (r *AlitripHotelHstdfShotelExportshotelRequest) SetTelNumber(telNumber string) error {
    r.telNumber = telNumber
    r.Set("tel_number", telNumber)
    return nil
}

// TelNumber Getter
func (r AlitripHotelHstdfShotelExportshotelRequest) GetTelNumber() string {
    return r.telNumber
}
// Longitude Setter
// 经度
func (r *AlitripHotelHstdfShotelExportshotelRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

// Longitude Getter
func (r AlitripHotelHstdfShotelExportshotelRequest) GetLongitude() string {
    return r.longitude
}
// Latitude Setter
// 纬度
func (r *AlitripHotelHstdfShotelExportshotelRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

// Latitude Getter
func (r AlitripHotelHstdfShotelExportshotelRequest) GetLatitude() string {
    return r.latitude
}
// Address Setter
// 地址
func (r *AlitripHotelHstdfShotelExportshotelRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

// Address Getter
func (r AlitripHotelHstdfShotelExportshotelRequest) GetAddress() string {
    return r.address
}
