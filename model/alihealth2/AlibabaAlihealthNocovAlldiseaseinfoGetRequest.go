package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取全国疫情统计数据 API请求
alibaba.alihealth.nocov.alldiseaseinfo.get

获取全国疫情统计数据
*/
type AlibabaAlihealthNocovAlldiseaseinfoGetRequest struct {
    model.Params
    // 省的
    province   string
    // 城市
    city   string
    // 城市code
    cityCode   string
}

// 初始化AlibabaAlihealthNocovAlldiseaseinfoGetRequest对象
func NewAlibabaAlihealthNocovAlldiseaseinfoGetRequest() *AlibabaAlihealthNocovAlldiseaseinfoGetRequest{
    return &AlibabaAlihealthNocovAlldiseaseinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthNocovAlldiseaseinfoGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.nocov.alldiseaseinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthNocovAlldiseaseinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Province Setter
// 省的
func (r *AlibabaAlihealthNocovAlldiseaseinfoGetRequest) SetProvince(province string) error {
    r.province = province
    r.Set("province", province)
    return nil
}

// Province Getter
func (r AlibabaAlihealthNocovAlldiseaseinfoGetRequest) GetProvince() string {
    return r.province
}
// City Setter
// 城市
func (r *AlibabaAlihealthNocovAlldiseaseinfoGetRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

// City Getter
func (r AlibabaAlihealthNocovAlldiseaseinfoGetRequest) GetCity() string {
    return r.city
}
// CityCode Setter
// 城市code
func (r *AlibabaAlihealthNocovAlldiseaseinfoGetRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

// CityCode Getter
func (r AlibabaAlihealthNocovAlldiseaseinfoGetRequest) GetCityCode() string {
    return r.cityCode
}
