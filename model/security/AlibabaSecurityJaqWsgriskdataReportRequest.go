package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
无线保镖SDK风控数据上报 API请求
alibaba.security.jaq.wsgriskdata.report

无线保镖sdk根据用户的需要，上报数据到聚安全云端
*/
type AlibabaSecurityJaqWsgriskdataReportRequest struct {
    model.Params
    // wua串
    _wua   string
    // mtopappkey是mtop的appkey
    _extParam   string
}

// 初始化AlibabaSecurityJaqWsgriskdataReportRequest对象
func NewAlibabaSecurityJaqWsgriskdataReportRequest() *AlibabaSecurityJaqWsgriskdataReportRequest{
    return &AlibabaSecurityJaqWsgriskdataReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqWsgriskdataReportRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.wsgriskdata.report"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqWsgriskdataReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Wua Setter
// wua串
func (r *AlibabaSecurityJaqWsgriskdataReportRequest) SetWua(_wua string) error {
    r._wua = _wua
    r.Set("wua", _wua)
    return nil
}

// Wua Getter
func (r AlibabaSecurityJaqWsgriskdataReportRequest) GetWua() string {
    return r._wua
}
// ExtParam Setter
// mtopappkey是mtop的appkey
func (r *AlibabaSecurityJaqWsgriskdataReportRequest) SetExtParam(_extParam string) error {
    r._extParam = _extParam
    r.Set("ext_param", _extParam)
    return nil
}

// ExtParam Getter
func (r AlibabaSecurityJaqWsgriskdataReportRequest) GetExtParam() string {
    return r._extParam
}
