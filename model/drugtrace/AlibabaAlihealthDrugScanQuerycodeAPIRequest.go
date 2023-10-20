package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugscanquerycodeAPIRequest 查询药监码对应的有效期和包装规格 API请求
// alibaba.alihealth.drug.scan.querycode
//
// 查询药监码对应的有效期和包装规格
type AlibabaalihealthdrugscanquerycodeAPIRequest struct {
	model.Params
	// 溯源码
	_code string
	// 扫码日期
	_scanTime string
	// 用户标识id
	_webchatId string
	// 省编码
	_provinceCode string
	// 市编码
	_cityCode string
	// 区编码
	_areaCode string
}

// NewAlibabaalihealthdrugscanquerycodeRequest 初始化AlibabaalihealthdrugscanquerycodeAPIRequest对象
func NewAlibabaalihealthdrugscanquerycodeRequest() *AlibabaalihealthdrugscanquerycodeAPIRequest {
	return &AlibabaalihealthdrugscanquerycodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugscanquerycodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.scan.querycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugscanquerycodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugscanquerycodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 溯源码
func (r *AlibabaalihealthdrugscanquerycodeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugscanquerycodeAPIRequest) GetCode() string {
	return r._code
}

// SetScanTime is ScanTime Setter
// 扫码日期
func (r *AlibabaalihealthdrugscanquerycodeAPIRequest) SetScanTime(_scanTime string) error {
	r._scanTime = _scanTime
	r.Set("scan_time", _scanTime)
	return nil
}

// GetScanTime ScanTime Getter
func (r AlibabaalihealthdrugscanquerycodeAPIRequest) GetScanTime() string {
	return r._scanTime
}

// SetWebchatId is WebchatId Setter
// 用户标识id
func (r *AlibabaalihealthdrugscanquerycodeAPIRequest) SetWebchatId(_webchatId string) error {
	r._webchatId = _webchatId
	r.Set("webchat_id", _webchatId)
	return nil
}

// GetWebchatId WebchatId Getter
func (r AlibabaalihealthdrugscanquerycodeAPIRequest) GetWebchatId() string {
	return r._webchatId
}

// SetProvinceCode is ProvinceCode Setter
// 省编码
func (r *AlibabaalihealthdrugscanquerycodeAPIRequest) SetProvinceCode(_provinceCode string) error {
	r._provinceCode = _provinceCode
	r.Set("province_code", _provinceCode)
	return nil
}

// GetProvinceCode ProvinceCode Getter
func (r AlibabaalihealthdrugscanquerycodeAPIRequest) GetProvinceCode() string {
	return r._provinceCode
}

// SetCityCode is CityCode Setter
// 市编码
func (r *AlibabaalihealthdrugscanquerycodeAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r AlibabaalihealthdrugscanquerycodeAPIRequest) GetCityCode() string {
	return r._cityCode
}

// SetAreaCode is AreaCode Setter
// 区编码
func (r *AlibabaalihealthdrugscanquerycodeAPIRequest) SetAreaCode(_areaCode string) error {
	r._areaCode = _areaCode
	r.Set("area_code", _areaCode)
	return nil
}

// GetAreaCode AreaCode Getter
func (r AlibabaalihealthdrugscanquerycodeAPIRequest) GetAreaCode() string {
	return r._areaCode
}
