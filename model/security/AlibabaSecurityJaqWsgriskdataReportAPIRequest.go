package security

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqWsgriskdataReportAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.wsgriskdata.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqWsgriskdataReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Wua Setter
// wua串
func (r *AlibabaSecurityJaqWsgriskdataReportAPIRequest) SetWua(_wua string) error {
	r._wua = _wua
	r.Set("wua", _wua)
	return nil
}

// Get Wua Getter
func (r AlibabaSecurityJaqWsgriskdataReportAPIRequest) GetWua() string {
	return r._wua
}

// Set is ExtParam Setter
// mtopappkey是mtop的appkey
func (r *AlibabaSecurityJaqWsgriskdataReportAPIRequest) SetExtParam(_extParam string) error {
	r._extParam = _extParam
	r.Set("ext_param", _extParam)
	return nil
}

// Get ExtParam Getter
func (r AlibabaSecurityJaqWsgriskdataReportAPIRequest) GetExtParam() string {
	return r._extParam
}
