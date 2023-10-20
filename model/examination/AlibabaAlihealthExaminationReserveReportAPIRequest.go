package examination

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReserveReportAPIRequest 体检机构对接_体检报告查询 API请求
// alibaba.alihealth.examination.reserve.report
//
// 体检机构对接_体检报告获取
type AlibabaAlihealthExaminationReserveReportAPIRequest struct {
	model.Params
	// 商户唯一码
	_merchantCode string
	// 阿里健康预约唯一标识
	_reserveNumber string
	// 到检唯一标识
	_checkNo string
	// 体检机构预约唯一标识码
	_uniqReserveCode string
	// 查询报告卡号
	_searchNo string
	// 查询报告密码
	_searchPwd string
}

// NewAlibabaAlihealthExaminationReserveReportRequest 初始化AlibabaAlihealthExaminationReserveReportAPIRequest对象
func NewAlibabaAlihealthExaminationReserveReportRequest() *AlibabaAlihealthExaminationReserveReportAPIRequest {
	return &AlibabaAlihealthExaminationReserveReportAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthExaminationReserveReportAPIRequest) Reset() {
	r._merchantCode = ""
	r._reserveNumber = ""
	r._checkNo = ""
	r._uniqReserveCode = ""
	r._searchNo = ""
	r._searchPwd = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReserveReportAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.reserve.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReserveReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthExaminationReserveReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMerchantCode is MerchantCode Setter
// 商户唯一码
func (r *AlibabaAlihealthExaminationReserveReportAPIRequest) SetMerchantCode(_merchantCode string) error {
	r._merchantCode = _merchantCode
	r.Set("merchant_code", _merchantCode)
	return nil
}

// GetMerchantCode MerchantCode Getter
func (r AlibabaAlihealthExaminationReserveReportAPIRequest) GetMerchantCode() string {
	return r._merchantCode
}

// SetReserveNumber is ReserveNumber Setter
// 阿里健康预约唯一标识
func (r *AlibabaAlihealthExaminationReserveReportAPIRequest) SetReserveNumber(_reserveNumber string) error {
	r._reserveNumber = _reserveNumber
	r.Set("reserve_number", _reserveNumber)
	return nil
}

// GetReserveNumber ReserveNumber Getter
func (r AlibabaAlihealthExaminationReserveReportAPIRequest) GetReserveNumber() string {
	return r._reserveNumber
}

// SetCheckNo is CheckNo Setter
// 到检唯一标识
func (r *AlibabaAlihealthExaminationReserveReportAPIRequest) SetCheckNo(_checkNo string) error {
	r._checkNo = _checkNo
	r.Set("check_no", _checkNo)
	return nil
}

// GetCheckNo CheckNo Getter
func (r AlibabaAlihealthExaminationReserveReportAPIRequest) GetCheckNo() string {
	return r._checkNo
}

// SetUniqReserveCode is UniqReserveCode Setter
// 体检机构预约唯一标识码
func (r *AlibabaAlihealthExaminationReserveReportAPIRequest) SetUniqReserveCode(_uniqReserveCode string) error {
	r._uniqReserveCode = _uniqReserveCode
	r.Set("uniq_reserve_code", _uniqReserveCode)
	return nil
}

// GetUniqReserveCode UniqReserveCode Getter
func (r AlibabaAlihealthExaminationReserveReportAPIRequest) GetUniqReserveCode() string {
	return r._uniqReserveCode
}

// SetSearchNo is SearchNo Setter
// 查询报告卡号
func (r *AlibabaAlihealthExaminationReserveReportAPIRequest) SetSearchNo(_searchNo string) error {
	r._searchNo = _searchNo
	r.Set("search_no", _searchNo)
	return nil
}

// GetSearchNo SearchNo Getter
func (r AlibabaAlihealthExaminationReserveReportAPIRequest) GetSearchNo() string {
	return r._searchNo
}

// SetSearchPwd is SearchPwd Setter
// 查询报告密码
func (r *AlibabaAlihealthExaminationReserveReportAPIRequest) SetSearchPwd(_searchPwd string) error {
	r._searchPwd = _searchPwd
	r.Set("search_pwd", _searchPwd)
	return nil
}

// GetSearchPwd SearchPwd Getter
func (r AlibabaAlihealthExaminationReserveReportAPIRequest) GetSearchPwd() string {
	return r._searchPwd
}

var poolAlibabaAlihealthExaminationReserveReportAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthExaminationReserveReportRequest()
	},
}

// GetAlibabaAlihealthExaminationReserveReportRequest 从 sync.Pool 获取 AlibabaAlihealthExaminationReserveReportAPIRequest
func GetAlibabaAlihealthExaminationReserveReportAPIRequest() *AlibabaAlihealthExaminationReserveReportAPIRequest {
	return poolAlibabaAlihealthExaminationReserveReportAPIRequest.Get().(*AlibabaAlihealthExaminationReserveReportAPIRequest)
}

// ReleaseAlibabaAlihealthExaminationReserveReportAPIRequest 将 AlibabaAlihealthExaminationReserveReportAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthExaminationReserveReportAPIRequest(v *AlibabaAlihealthExaminationReserveReportAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthExaminationReserveReportAPIRequest.Put(v)
}
