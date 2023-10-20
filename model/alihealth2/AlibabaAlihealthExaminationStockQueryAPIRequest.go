package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationstockqueryAPIRequest 体检机构对接_体检套餐库存查询 API请求
// alibaba.alihealth.examination.stock.query
//
// 体检机构对接_体检套餐库存查询
type AlibabaalihealthexaminationstockqueryAPIRequest struct {
	model.Params
	// 商户唯一码
	_merchantCode string
	// 分院ID
	_hospitalId string
	// 套餐id
	_packageId string
	// 开始日期
	_timeFrom string
	// 结束日期
	_timeTo string
}

// NewAlibabaalihealthexaminationstockqueryRequest 初始化AlibabaalihealthexaminationstockqueryAPIRequest对象
func NewAlibabaalihealthexaminationstockqueryRequest() *AlibabaalihealthexaminationstockqueryAPIRequest {
	return &AlibabaalihealthexaminationstockqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationstockqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.stock.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationstockqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationstockqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMerchantCode is MerchantCode Setter
// 商户唯一码
func (r *AlibabaalihealthexaminationstockqueryAPIRequest) SetMerchantCode(_merchantCode string) error {
	r._merchantCode = _merchantCode
	r.Set("merchant_code", _merchantCode)
	return nil
}

// GetMerchantCode MerchantCode Getter
func (r AlibabaalihealthexaminationstockqueryAPIRequest) GetMerchantCode() string {
	return r._merchantCode
}

// SetHospitalId is HospitalId Setter
// 分院ID
func (r *AlibabaalihealthexaminationstockqueryAPIRequest) SetHospitalId(_hospitalId string) error {
	r._hospitalId = _hospitalId
	r.Set("hospital_id", _hospitalId)
	return nil
}

// GetHospitalId HospitalId Getter
func (r AlibabaalihealthexaminationstockqueryAPIRequest) GetHospitalId() string {
	return r._hospitalId
}

// SetPackageId is PackageId Setter
// 套餐id
func (r *AlibabaalihealthexaminationstockqueryAPIRequest) SetPackageId(_packageId string) error {
	r._packageId = _packageId
	r.Set("package_id", _packageId)
	return nil
}

// GetPackageId PackageId Getter
func (r AlibabaalihealthexaminationstockqueryAPIRequest) GetPackageId() string {
	return r._packageId
}

// SetTimeFrom is TimeFrom Setter
// 开始日期
func (r *AlibabaalihealthexaminationstockqueryAPIRequest) SetTimeFrom(_timeFrom string) error {
	r._timeFrom = _timeFrom
	r.Set("time_from", _timeFrom)
	return nil
}

// GetTimeFrom TimeFrom Getter
func (r AlibabaalihealthexaminationstockqueryAPIRequest) GetTimeFrom() string {
	return r._timeFrom
}

// SetTimeTo is TimeTo Setter
// 结束日期
func (r *AlibabaalihealthexaminationstockqueryAPIRequest) SetTimeTo(_timeTo string) error {
	r._timeTo = _timeTo
	r.Set("time_to", _timeTo)
	return nil
}

// GetTimeTo TimeTo Getter
func (r AlibabaalihealthexaminationstockqueryAPIRequest) GetTimeTo() string {
	return r._timeTo
}
