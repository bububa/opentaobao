package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugScanQuerycodeAPIRequest 查询药监码对应的有效期和包装规格 API请求
// alibaba.alihealth.drug.scan.querycode
//
// 查询药监码对应的有效期和包装规格
type AlibabaAlihealthDrugScanQuerycodeAPIRequest struct {
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

// NewAlibabaAlihealthDrugScanQuerycodeRequest 初始化AlibabaAlihealthDrugScanQuerycodeAPIRequest对象
func NewAlibabaAlihealthDrugScanQuerycodeRequest() *AlibabaAlihealthDrugScanQuerycodeAPIRequest {
	return &AlibabaAlihealthDrugScanQuerycodeAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugScanQuerycodeAPIRequest) Reset() {
	r._code = ""
	r._scanTime = ""
	r._webchatId = ""
	r._provinceCode = ""
	r._cityCode = ""
	r._areaCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugScanQuerycodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.scan.querycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugScanQuerycodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugScanQuerycodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 溯源码
func (r *AlibabaAlihealthDrugScanQuerycodeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugScanQuerycodeAPIRequest) GetCode() string {
	return r._code
}

// SetScanTime is ScanTime Setter
// 扫码日期
func (r *AlibabaAlihealthDrugScanQuerycodeAPIRequest) SetScanTime(_scanTime string) error {
	r._scanTime = _scanTime
	r.Set("scan_time", _scanTime)
	return nil
}

// GetScanTime ScanTime Getter
func (r AlibabaAlihealthDrugScanQuerycodeAPIRequest) GetScanTime() string {
	return r._scanTime
}

// SetWebchatId is WebchatId Setter
// 用户标识id
func (r *AlibabaAlihealthDrugScanQuerycodeAPIRequest) SetWebchatId(_webchatId string) error {
	r._webchatId = _webchatId
	r.Set("webchat_id", _webchatId)
	return nil
}

// GetWebchatId WebchatId Getter
func (r AlibabaAlihealthDrugScanQuerycodeAPIRequest) GetWebchatId() string {
	return r._webchatId
}

// SetProvinceCode is ProvinceCode Setter
// 省编码
func (r *AlibabaAlihealthDrugScanQuerycodeAPIRequest) SetProvinceCode(_provinceCode string) error {
	r._provinceCode = _provinceCode
	r.Set("province_code", _provinceCode)
	return nil
}

// GetProvinceCode ProvinceCode Getter
func (r AlibabaAlihealthDrugScanQuerycodeAPIRequest) GetProvinceCode() string {
	return r._provinceCode
}

// SetCityCode is CityCode Setter
// 市编码
func (r *AlibabaAlihealthDrugScanQuerycodeAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r AlibabaAlihealthDrugScanQuerycodeAPIRequest) GetCityCode() string {
	return r._cityCode
}

// SetAreaCode is AreaCode Setter
// 区编码
func (r *AlibabaAlihealthDrugScanQuerycodeAPIRequest) SetAreaCode(_areaCode string) error {
	r._areaCode = _areaCode
	r.Set("area_code", _areaCode)
	return nil
}

// GetAreaCode AreaCode Getter
func (r AlibabaAlihealthDrugScanQuerycodeAPIRequest) GetAreaCode() string {
	return r._areaCode
}

var poolAlibabaAlihealthDrugScanQuerycodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugScanQuerycodeRequest()
	},
}

// GetAlibabaAlihealthDrugScanQuerycodeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugScanQuerycodeAPIRequest
func GetAlibabaAlihealthDrugScanQuerycodeAPIRequest() *AlibabaAlihealthDrugScanQuerycodeAPIRequest {
	return poolAlibabaAlihealthDrugScanQuerycodeAPIRequest.Get().(*AlibabaAlihealthDrugScanQuerycodeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugScanQuerycodeAPIRequest 将 AlibabaAlihealthDrugScanQuerycodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugScanQuerycodeAPIRequest(v *AlibabaAlihealthDrugScanQuerycodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugScanQuerycodeAPIRequest.Put(v)
}
