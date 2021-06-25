package dt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线索报价价格校验 APIRequest
alibaba.dt.tmllcar.pricevalidate

根据选定的车型和城市，校验汽车价格是否通过
入参：车型ID，城市名称，价格
输出：N 校验失败，校验成功不返回值
*/
type AlibabaDtTmllcarPricevalidateRequest struct {
    model.Params

    // tt
    appName   string 

    // tt
    name   string 

    // tt
    password   string 

    // tt
    price   *BigDecimal 

    // tt
    cityName   string 

    // tt
    modelName   string 

}

func NewAlibabaDtTmllcarPricevalidateRequest() *AlibabaDtTmllcarPricevalidateRequest{
    return &AlibabaDtTmllcarPricevalidateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDtTmllcarPricevalidateRequest) GetApiMethodName() string {
    return "alibaba.dt.tmllcar.pricevalidate"
}

func (r AlibabaDtTmllcarPricevalidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDtTmllcarPricevalidateRequest) SetAppName(appName string) error {
    r.appName = appName
    r.Set("app_name", appName)
    return nil
}

func (r AlibabaDtTmllcarPricevalidateRequest) GetAppName() string {
    return r.appName
}

func (r *AlibabaDtTmllcarPricevalidateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r AlibabaDtTmllcarPricevalidateRequest) GetName() string {
    return r.name
}

func (r *AlibabaDtTmllcarPricevalidateRequest) SetPassword(password string) error {
    r.password = password
    r.Set("password", password)
    return nil
}

func (r AlibabaDtTmllcarPricevalidateRequest) GetPassword() string {
    return r.password
}

func (r *AlibabaDtTmllcarPricevalidateRequest) SetPrice(price *BigDecimal) error {
    r.price = price
    r.Set("price", price)
    return nil
}

func (r AlibabaDtTmllcarPricevalidateRequest) GetPrice() *BigDecimal {
    return r.price
}

func (r *AlibabaDtTmllcarPricevalidateRequest) SetCityName(cityName string) error {
    r.cityName = cityName
    r.Set("city_name", cityName)
    return nil
}

func (r AlibabaDtTmllcarPricevalidateRequest) GetCityName() string {
    return r.cityName
}

func (r *AlibabaDtTmllcarPricevalidateRequest) SetModelName(modelName string) error {
    r.modelName = modelName
    r.Set("model_name", modelName)
    return nil
}

func (r AlibabaDtTmllcarPricevalidateRequest) GetModelName() string {
    return r.modelName
}

