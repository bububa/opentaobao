package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取全国疫情统计数据 APIRequest
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

func NewAlibabaAlihealthNocovAlldiseaseinfoGetRequest() *AlibabaAlihealthNocovAlldiseaseinfoGetRequest{
    return &AlibabaAlihealthNocovAlldiseaseinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthNocovAlldiseaseinfoGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.nocov.alldiseaseinfo.get"
}

func (r AlibabaAlihealthNocovAlldiseaseinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthNocovAlldiseaseinfoGetRequest) SetProvince(province string) error {
    r.province = province
    r.Set("province", province)
    return nil
}

func (r AlibabaAlihealthNocovAlldiseaseinfoGetRequest) GetProvince() string {
    return r.province
}

func (r *AlibabaAlihealthNocovAlldiseaseinfoGetRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

func (r AlibabaAlihealthNocovAlldiseaseinfoGetRequest) GetCity() string {
    return r.city
}

func (r *AlibabaAlihealthNocovAlldiseaseinfoGetRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

func (r AlibabaAlihealthNocovAlldiseaseinfoGetRequest) GetCityCode() string {
    return r.cityCode
}

