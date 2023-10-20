package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest 数据未落地诊断 API请求
// alibaba.alihealth.drug.download.dataerrordiagnosis
//
// 阿里健康-追溯码-D2D数据未落地原因诊断
type AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest struct {
	model.Params
	// appKey
	_appKeyN string
	// 数据所有者企业名称
	_baseEntName string
	// 单据所有者企业名称
	_billEntName string
	// 下游模式填2 集团模式填3
	_type string
	// 单据号
	_billCode string
	// 单据标识 入库填写I 出库填写O
	_billTypeFlag string
	// 是否需要重传 1代表需要 0代表不需要
	_reUpload string
	// 追溯码；当有code时候billEntname  bill_code  bill_type_flag可以不填，优先根据code判定
	_code string
}

// NewAlibabaalihealthdrugdownloaddataerrordiagnosisRequest 初始化AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest对象
func NewAlibabaalihealthdrugdownloaddataerrordiagnosisRequest() *AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest {
	return &AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.download.dataerrordiagnosis"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppKeyN is AppKeyN Setter
// appKey
func (r *AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) SetAppKeyN(_appKeyN string) error {
	r._appKeyN = _appKeyN
	r.Set("app_key_n", _appKeyN)
	return nil
}

// GetAppKeyN AppKeyN Getter
func (r AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) GetAppKeyN() string {
	return r._appKeyN
}

// SetBaseEntName is BaseEntName Setter
// 数据所有者企业名称
func (r *AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) SetBaseEntName(_baseEntName string) error {
	r._baseEntName = _baseEntName
	r.Set("base_ent_name", _baseEntName)
	return nil
}

// GetBaseEntName BaseEntName Getter
func (r AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) GetBaseEntName() string {
	return r._baseEntName
}

// SetBillEntName is BillEntName Setter
// 单据所有者企业名称
func (r *AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) SetBillEntName(_billEntName string) error {
	r._billEntName = _billEntName
	r.Set("bill_ent_name", _billEntName)
	return nil
}

// GetBillEntName BillEntName Getter
func (r AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) GetBillEntName() string {
	return r._billEntName
}

// SetType is Type Setter
// 下游模式填2 集团模式填3
func (r *AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) GetType() string {
	return r._type
}

// SetBillCode is BillCode Setter
// 单据号
func (r *AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillTypeFlag is BillTypeFlag Setter
// 单据标识 入库填写I 出库填写O
func (r *AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) SetBillTypeFlag(_billTypeFlag string) error {
	r._billTypeFlag = _billTypeFlag
	r.Set("bill_type_flag", _billTypeFlag)
	return nil
}

// GetBillTypeFlag BillTypeFlag Getter
func (r AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) GetBillTypeFlag() string {
	return r._billTypeFlag
}

// SetReUpload is ReUpload Setter
// 是否需要重传 1代表需要 0代表不需要
func (r *AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) SetReUpload(_reUpload string) error {
	r._reUpload = _reUpload
	r.Set("re_upload", _reUpload)
	return nil
}

// GetReUpload ReUpload Getter
func (r AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) GetReUpload() string {
	return r._reUpload
}

// SetCode is Code Setter
// 追溯码；当有code时候billEntname  bill_code  bill_type_flag可以不填，优先根据code判定
func (r *AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugdownloaddataerrordiagnosisAPIRequest) GetCode() string {
	return r._code
}
