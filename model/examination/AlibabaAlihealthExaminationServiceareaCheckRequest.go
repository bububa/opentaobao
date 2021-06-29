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
    packageCode   string
    // 上门检测地址
    address   string
    // 上门检测地址纬度
    latitude   string
    // 上门检测地址经度
    longitude   string
    // 省份名称（高德）
    province   string
    // 省份编码（高德adcode）
    provinceCode   string
    // 城市名称（高德）
    city   string
    // 城市编码（高德adcode）
    cityCode   string
    // 区域名称（高德）
    district   string
    // 区域编码（高德adcode）
    districtCode   string
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
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetPackageCode(packageCode string) error {
    r.packageCode = packageCode
    r.Set("package_code", packageCode)
    return nil
}

// PackageCode Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetPackageCode() string {
    return r.packageCode
}
// Address Setter
// 上门检测地址
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

// Address Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetAddress() string {
    return r.address
}
// Latitude Setter
// 上门检测地址纬度
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

// Latitude Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetLatitude() string {
    return r.latitude
}
// Longitude Setter
// 上门检测地址经度
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

// Longitude Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetLongitude() string {
    return r.longitude
}
// Province Setter
// 省份名称（高德）
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetProvince(province string) error {
    r.province = province
    r.Set("province", province)
    return nil
}

// Province Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetProvince() string {
    return r.province
}
// ProvinceCode Setter
// 省份编码（高德adcode）
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetProvinceCode(provinceCode string) error {
    r.provinceCode = provinceCode
    r.Set("province_code", provinceCode)
    return nil
}

// ProvinceCode Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetProvinceCode() string {
    return r.provinceCode
}
// City Setter
// 城市名称（高德）
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

// City Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetCity() string {
    return r.city
}
// CityCode Setter
// 城市编码（高德adcode）
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

// CityCode Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetCityCode() string {
    return r.cityCode
}
// District Setter
// 区域名称（高德）
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetDistrict(district string) error {
    r.district = district
    r.Set("district", district)
    return nil
}

// District Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetDistrict() string {
    return r.district
}
// DistrictCode Setter
// 区域编码（高德adcode）
func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetDistrictCode(districtCode string) error {
    r.districtCode = districtCode
    r.Set("district_code", districtCode)
    return nil
}

// DistrictCode Getter
func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetDistrictCode() string {
    return r.districtCode
}
