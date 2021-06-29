package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构对接_上门检测服务范围查询 API请求
alibaba.alihealth.examination.servicearea.check

体检机构对接_上门检测服务范围查询
*/
type AlibabaAlihealthExaminationServiceareaCheckRequest struct {
    model.Params
    // 机构套餐编码
    _packageCode   string
    // 上门检测地址
    _address   string
    // 上门检测地址纬度
    _latitude   string
    // 上门检测地址经度
    _longitude   string
    // 省份名称（高德）
    _province   string
    // 省份编码（高德adcode）
    _provinceCode   string
    // 城市名称（高德）
    _city   string
    // 城市编码（高德adcode）
    _cityCode   string
    // 区域名称（高德）
    _district   string
    // 区域编码（高德adcode）
    _districtCode   string
}

// 初始化AlibabaAlihealthExaminationServiceareaCheckRequest对象
func NewAlibabaAlihealthExaminationServiceareaCheckRequest() *AlibabaAlihealthExaminationServiceareaCheckRequest{
    return &AlibabaAlihealthExaminationServiceareaCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.servicearea.check"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PackageCode Setter
// 机构套餐编码
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetPackageCode(_packageCode string) error {
    r._packageCode = _packageCode
    r.Set("package_code", _packageCode)
    return nil
}

// PackageCode Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetPackageCode() string {
    return r._packageCode
}
// Address Setter
// 上门检测地址
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetAddress() string {
    return r._address
}
// Latitude Setter
// 上门检测地址纬度
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetLatitude(_latitude string) error {
    r._latitude = _latitude
    r.Set("latitude", _latitude)
    return nil
}

// Latitude Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetLatitude() string {
    return r._latitude
}
// Longitude Setter
// 上门检测地址经度
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetLongitude(_longitude string) error {
    r._longitude = _longitude
    r.Set("longitude", _longitude)
    return nil
}

// Longitude Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetLongitude() string {
    return r._longitude
}
// Province Setter
// 省份名称（高德）
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetProvince(_province string) error {
    r._province = _province
    r.Set("province", _province)
    return nil
}

// Province Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetProvince() string {
    return r._province
}
// ProvinceCode Setter
// 省份编码（高德adcode）
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetProvinceCode(_provinceCode string) error {
    r._provinceCode = _provinceCode
    r.Set("province_code", _provinceCode)
    return nil
}

// ProvinceCode Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetProvinceCode() string {
    return r._provinceCode
}
// City Setter
// 城市名称（高德）
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetCity(_city string) error {
    r._city = _city
    r.Set("city", _city)
    return nil
}

// City Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetCity() string {
    return r._city
}
// CityCode Setter
// 城市编码（高德adcode）
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetCityCode(_cityCode string) error {
    r._cityCode = _cityCode
    r.Set("city_code", _cityCode)
    return nil
}

// CityCode Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetCityCode() string {
    return r._cityCode
}
// District Setter
// 区域名称（高德）
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetDistrict(_district string) error {
    r._district = _district
    r.Set("district", _district)
    return nil
}

// District Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetDistrict() string {
    return r._district
}
// DistrictCode Setter
// 区域编码（高德adcode）
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetDistrictCode(_districtCode string) error {
    r._districtCode = _districtCode
    r.Set("district_code", _districtCode)
    return nil
}

// DistrictCode Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetDistrictCode() string {
    return r._districtCode
}
