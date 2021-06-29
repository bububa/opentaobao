package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构对接_上门检测服务范围查询 APIRequest
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

func NewAlibabaAlihealthExaminationServiceareaCheckRequest() *AlibabaAlihealthExaminationServiceareaCheckRequest{
    return &AlibabaAlihealthExaminationServiceareaCheckRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.servicearea.check"
}

func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetPackageCode(packageCode string) error {
    r.packageCode = packageCode
    r.Set("package_code", packageCode)
    return nil
}

func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetPackageCode() string {
    return r.packageCode
}

func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetAddress() string {
    return r.address
}

func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetLatitude() string {
    return r.latitude
}

func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetLongitude() string {
    return r.longitude
}

func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetProvince(province string) error {
    r.province = province
    r.Set("province", province)
    return nil
}

func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetProvince() string {
    return r.province
}

func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetProvinceCode(provinceCode string) error {
    r.provinceCode = provinceCode
    r.Set("province_code", provinceCode)
    return nil
}

func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetProvinceCode() string {
    return r.provinceCode
}

func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetCity() string {
    return r.city
}

func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetCityCode() string {
    return r.cityCode
}

func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetDistrict(district string) error {
    r.district = district
    r.Set("district", district)
    return nil
}

func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetDistrict() string {
    return r.district
}

func (r *AlibabaAlihealthExaminationServiceareaCheckRequest) SetDistrictCode(districtCode string) error {
    r.districtCode = districtCode
    r.Set("district_code", districtCode)
    return nil
}

func (r AlibabaAlihealthExaminationServiceareaCheckRequest) GetDistrictCode() string {
    return r.districtCode
}

