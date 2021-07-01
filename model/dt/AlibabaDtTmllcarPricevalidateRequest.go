package dt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线索报价价格校验 API请求
alibaba.dt.tmllcar.pricevalidate

根据选定的车型和城市，校验汽车价格是否通过
入参：车型ID，城市名称，价格
输出：N 校验失败，校验成功不返回值
*/
type AlibabaDtTmllcarPricevalidateAPIRequest struct {
    model.Params
    // tt
    _appName   string
    // tt
    _name   string
    // tt
    _password   string
    // tt
    _price   *BigDecimal
    // tt
    _cityName   string
    // tt
    _modelName   string
}

// 初始化AlibabaDtTmllcarPricevalidateAPIRequest对象
func NewAlibabaDtTmllcarPricevalidateRequest() *AlibabaDtTmllcarPricevalidateAPIRequest{
    return &AlibabaDtTmllcarPricevalidateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDtTmllcarPricevalidateAPIRequest) GetApiMethodName() string {
    return "alibaba.dt.tmllcar.pricevalidate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDtTmllcarPricevalidateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppName Setter
// tt
func (r *AlibabaDtTmllcarPricevalidateAPIRequest) SetAppName(_appName string) error {
    r._appName = _appName
    r.Set("app_name", _appName)
    return nil
}

// AppName Getter
func (r AlibabaDtTmllcarPricevalidateAPIRequest) GetAppName() string {
    return r._appName
}
// Name Setter
// tt
func (r *AlibabaDtTmllcarPricevalidateAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r AlibabaDtTmllcarPricevalidateAPIRequest) GetName() string {
    return r._name
}
// Password Setter
// tt
func (r *AlibabaDtTmllcarPricevalidateAPIRequest) SetPassword(_password string) error {
    r._password = _password
    r.Set("password", _password)
    return nil
}

// Password Getter
func (r AlibabaDtTmllcarPricevalidateAPIRequest) GetPassword() string {
    return r._password
}
// Price Setter
// tt
func (r *AlibabaDtTmllcarPricevalidateAPIRequest) SetPrice(_price *BigDecimal) error {
    r._price = _price
    r.Set("price", _price)
    return nil
}

// Price Getter
func (r AlibabaDtTmllcarPricevalidateAPIRequest) GetPrice() *BigDecimal {
    return r._price
}
// CityName Setter
// tt
func (r *AlibabaDtTmllcarPricevalidateAPIRequest) SetCityName(_cityName string) error {
    r._cityName = _cityName
    r.Set("city_name", _cityName)
    return nil
}

// CityName Getter
func (r AlibabaDtTmllcarPricevalidateAPIRequest) GetCityName() string {
    return r._cityName
}
// ModelName Setter
// tt
func (r *AlibabaDtTmllcarPricevalidateAPIRequest) SetModelName(_modelName string) error {
    r._modelName = _modelName
    r.Set("model_name", _modelName)
    return nil
}

// ModelName Getter
func (r AlibabaDtTmllcarPricevalidateAPIRequest) GetModelName() string {
    return r._modelName
}
