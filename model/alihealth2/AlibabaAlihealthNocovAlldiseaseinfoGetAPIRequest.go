package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthnocovalldiseaseinfogetAPIRequest 获取全国疫情统计数据 API请求
// alibaba.alihealth.nocov.alldiseaseinfo.get
//
// 获取全国疫情统计数据
type AlibabaalihealthnocovalldiseaseinfogetAPIRequest struct {
	model.Params
	// 省的
	_province string
	// 城市
	_city string
	// 城市code
	_cityCode string
}

// NewAlibabaalihealthnocovalldiseaseinfogetRequest 初始化AlibabaalihealthnocovalldiseaseinfogetAPIRequest对象
func NewAlibabaalihealthnocovalldiseaseinfogetRequest() *AlibabaalihealthnocovalldiseaseinfogetAPIRequest {
	return &AlibabaalihealthnocovalldiseaseinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthnocovalldiseaseinfogetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.nocov.alldiseaseinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthnocovalldiseaseinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthnocovalldiseaseinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProvince is Province Setter
// 省的
func (r *AlibabaalihealthnocovalldiseaseinfogetAPIRequest) SetProvince(_province string) error {
	r._province = _province
	r.Set("province", _province)
	return nil
}

// GetProvince Province Getter
func (r AlibabaalihealthnocovalldiseaseinfogetAPIRequest) GetProvince() string {
	return r._province
}

// SetCity is City Setter
// 城市
func (r *AlibabaalihealthnocovalldiseaseinfogetAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r AlibabaalihealthnocovalldiseaseinfogetAPIRequest) GetCity() string {
	return r._city
}

// SetCityCode is CityCode Setter
// 城市code
func (r *AlibabaalihealthnocovalldiseaseinfogetAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r AlibabaalihealthnocovalldiseaseinfogetAPIRequest) GetCityCode() string {
	return r._cityCode
}
