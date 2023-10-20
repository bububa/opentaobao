package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqWsgriskdataReportAPIRequest 无线保镖SDK风控数据上报 API请求
// alibaba.security.jaq.wsgriskdata.report
//
// 无线保镖sdk根据用户的需要，上报数据到聚安全云端
type AlibabaSecurityJaqWsgriskdataReportAPIRequest struct {
	model.Params
	// wua串
	_wua string
	// mtopappkey是mtop的appkey
	_extParam string
}

// NewAlibabaSecurityJaqWsgriskdataReportRequest 初始化AlibabaSecurityJaqWsgriskdataReportAPIRequest对象
func NewAlibabaSecurityJaqWsgriskdataReportRequest() *AlibabaSecurityJaqWsgriskdataReportAPIRequest {
	return &AlibabaSecurityJaqWsgriskdataReportAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqWsgriskdataReportAPIRequest) Reset() {
	r._wua = ""
	r._extParam = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqWsgriskdataReportAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.wsgriskdata.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqWsgriskdataReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqWsgriskdataReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWua is Wua Setter
// wua串
func (r *AlibabaSecurityJaqWsgriskdataReportAPIRequest) SetWua(_wua string) error {
	r._wua = _wua
	r.Set("wua", _wua)
	return nil
}

// GetWua Wua Getter
func (r AlibabaSecurityJaqWsgriskdataReportAPIRequest) GetWua() string {
	return r._wua
}

// SetExtParam is ExtParam Setter
// mtopappkey是mtop的appkey
func (r *AlibabaSecurityJaqWsgriskdataReportAPIRequest) SetExtParam(_extParam string) error {
	r._extParam = _extParam
	r.Set("ext_param", _extParam)
	return nil
}

// GetExtParam ExtParam Getter
func (r AlibabaSecurityJaqWsgriskdataReportAPIRequest) GetExtParam() string {
	return r._extParam
}

var poolAlibabaSecurityJaqWsgriskdataReportAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqWsgriskdataReportRequest()
	},
}

// GetAlibabaSecurityJaqWsgriskdataReportRequest 从 sync.Pool 获取 AlibabaSecurityJaqWsgriskdataReportAPIRequest
func GetAlibabaSecurityJaqWsgriskdataReportAPIRequest() *AlibabaSecurityJaqWsgriskdataReportAPIRequest {
	return poolAlibabaSecurityJaqWsgriskdataReportAPIRequest.Get().(*AlibabaSecurityJaqWsgriskdataReportAPIRequest)
}

// ReleaseAlibabaSecurityJaqWsgriskdataReportAPIRequest 将 AlibabaSecurityJaqWsgriskdataReportAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqWsgriskdataReportAPIRequest(v *AlibabaSecurityJaqWsgriskdataReportAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqWsgriskdataReportAPIRequest.Put(v)
}
