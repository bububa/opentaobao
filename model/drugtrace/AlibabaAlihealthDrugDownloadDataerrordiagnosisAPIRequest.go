package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest 数据未落地诊断 API请求
// alibaba.alihealth.drug.download.dataerrordiagnosis
//
// 阿里健康-追溯码-D2D数据未落地原因诊断
type AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest struct {
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

// NewAlibabaAlihealthDrugDownloadDataerrordiagnosisRequest 初始化AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest对象
func NewAlibabaAlihealthDrugDownloadDataerrordiagnosisRequest() *AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest {
	return &AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) Reset() {
	r._appKeyN = ""
	r._baseEntName = ""
	r._billEntName = ""
	r._type = ""
	r._billCode = ""
	r._billTypeFlag = ""
	r._reUpload = ""
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.download.dataerrordiagnosis"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppKeyN is AppKeyN Setter
// appKey
func (r *AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) SetAppKeyN(_appKeyN string) error {
	r._appKeyN = _appKeyN
	r.Set("app_key_n", _appKeyN)
	return nil
}

// GetAppKeyN AppKeyN Getter
func (r AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) GetAppKeyN() string {
	return r._appKeyN
}

// SetBaseEntName is BaseEntName Setter
// 数据所有者企业名称
func (r *AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) SetBaseEntName(_baseEntName string) error {
	r._baseEntName = _baseEntName
	r.Set("base_ent_name", _baseEntName)
	return nil
}

// GetBaseEntName BaseEntName Getter
func (r AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) GetBaseEntName() string {
	return r._baseEntName
}

// SetBillEntName is BillEntName Setter
// 单据所有者企业名称
func (r *AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) SetBillEntName(_billEntName string) error {
	r._billEntName = _billEntName
	r.Set("bill_ent_name", _billEntName)
	return nil
}

// GetBillEntName BillEntName Getter
func (r AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) GetBillEntName() string {
	return r._billEntName
}

// SetType is Type Setter
// 下游模式填2 集团模式填3
func (r *AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) GetType() string {
	return r._type
}

// SetBillCode is BillCode Setter
// 单据号
func (r *AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillTypeFlag is BillTypeFlag Setter
// 单据标识 入库填写I 出库填写O
func (r *AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) SetBillTypeFlag(_billTypeFlag string) error {
	r._billTypeFlag = _billTypeFlag
	r.Set("bill_type_flag", _billTypeFlag)
	return nil
}

// GetBillTypeFlag BillTypeFlag Getter
func (r AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) GetBillTypeFlag() string {
	return r._billTypeFlag
}

// SetReUpload is ReUpload Setter
// 是否需要重传 1代表需要 0代表不需要
func (r *AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) SetReUpload(_reUpload string) error {
	r._reUpload = _reUpload
	r.Set("re_upload", _reUpload)
	return nil
}

// GetReUpload ReUpload Getter
func (r AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) GetReUpload() string {
	return r._reUpload
}

// SetCode is Code Setter
// 追溯码；当有code时候billEntname  bill_code  bill_type_flag可以不填，优先根据code判定
func (r *AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) GetCode() string {
	return r._code
}

var poolAlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugDownloadDataerrordiagnosisRequest()
	},
}

// GetAlibabaAlihealthDrugDownloadDataerrordiagnosisRequest 从 sync.Pool 获取 AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest
func GetAlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest() *AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest {
	return poolAlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest.Get().(*AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest)
}

// ReleaseAlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest 将 AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest(v *AlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugDownloadDataerrordiagnosisAPIRequest.Put(v)
}
