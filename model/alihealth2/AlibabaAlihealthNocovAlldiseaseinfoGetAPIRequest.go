package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest 获取全国疫情统计数据 API请求
// alibaba.alihealth.nocov.alldiseaseinfo.get
//
// 获取全国疫情统计数据
type AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest struct {
	model.Params
	// 省的
	_province string
	// 城市
	_city string
	// 城市code
	_cityCode string
}

// NewAlibabaAlihealthNocovAlldiseaseinfoGetRequest 初始化AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest对象
func NewAlibabaAlihealthNocovAlldiseaseinfoGetRequest() *AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest {
	return &AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) Reset() {
	r._province = ""
	r._city = ""
	r._cityCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.nocov.alldiseaseinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProvince is Province Setter
// 省的
func (r *AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) SetProvince(_province string) error {
	r._province = _province
	r.Set("province", _province)
	return nil
}

// GetProvince Province Getter
func (r AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) GetProvince() string {
	return r._province
}

// SetCity is City Setter
// 城市
func (r *AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) GetCity() string {
	return r._city
}

// SetCityCode is CityCode Setter
// 城市code
func (r *AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) GetCityCode() string {
	return r._cityCode
}

var poolAlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthNocovAlldiseaseinfoGetRequest()
	},
}

// GetAlibabaAlihealthNocovAlldiseaseinfoGetRequest 从 sync.Pool 获取 AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest
func GetAlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest() *AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest {
	return poolAlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest.Get().(*AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest)
}

// ReleaseAlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest 将 AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest(v *AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest.Put(v)
}
