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
    _hid   int64
    // 酒店名称，必填
    _name   string
    // 酒店所在行政区划，对应平台ID，为空时会使用经纬度来定位
    _cityCode   int64
    // 电话
    _telNumber   string
    // 经度
    _longitude   string
    // 纬度
    _latitude   string
    // 地址
    _address   string
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
func (r *AlitripHotelHstdfShotelExportshotelRequest) SetHid(_hid int64) error {
    r._hid = _hid
    r.Set("hid", _hid)
    return nil
}

// Hid Getter
func (r AlitripHotelHstdfShotelExportshotelRequest) GetHid() int64 {
    return r._hid
}
// Name Setter
// 酒店名称，必填
func (r *AlitripHotelHstdfShotelExportshotelRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r AlitripHotelHstdfShotelExportshotelRequest) GetName() string {
    return r._name
}
// CityCode Setter
// 酒店所在行政区划，对应平台ID，为空时会使用经纬度来定位
func (r *AlitripHotelHstdfShotelExportshotelRequest) SetCityCode(_cityCode int64) error {
    r._cityCode = _cityCode
    r.Set("city_code", _cityCode)
    return nil
}

// CityCode Getter
func (r AlitripHotelHstdfShotelExportshotelRequest) GetCityCode() int64 {
    return r._cityCode
}
// TelNumber Setter
// 电话
func (r *AlitripHotelHstdfShotelExportshotelRequest) SetTelNumber(_telNumber string) error {
    r._telNumber = _telNumber
    r.Set("tel_number", _telNumber)
    return nil
}

// TelNumber Getter
func (r AlitripHotelHstdfShotelExportshotelRequest) GetTelNumber() string {
    return r._telNumber
}
// Longitude Setter
// 经度
func (r *AlitripHotelHstdfShotelExportshotelRequest) SetLongitude(_longitude string) error {
    r._longitude = _longitude
    r.Set("longitude", _longitude)
    return nil
}

// Longitude Getter
func (r AlitripHotelHstdfShotelExportshotelRequest) GetLongitude() string {
    return r._longitude
}
// Latitude Setter
// 纬度
func (r *AlitripHotelHstdfShotelExportshotelRequest) SetLatitude(_latitude string) error {
    r._latitude = _latitude
    r.Set("latitude", _latitude)
    return nil
}

// Latitude Getter
func (r AlitripHotelHstdfShotelExportshotelRequest) GetLatitude() string {
    return r._latitude
}
// Address Setter
// 地址
func (r *AlitripHotelHstdfShotelExportshotelRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r AlitripHotelHstdfShotelExportshotelRequest) GetAddress() string {
    return r._address
}
