package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
无线保镖SDK风控数据上报 APIRequest
alibaba.security.jaq.wsgriskdata.report

无线保镖sdk根据用户的需要，上报数据到聚安全云端
*/
type AlibabaSecurityJaqWsgriskdataReportRequest struct {
    model.Params

    // wua串
    wua   string 

    // mtopappkey是mtop的appkey
    extParam   string 

}

func NewAlibabaSecurityJaqWsgriskdataReportRequest() *AlibabaSecurityJaqWsgriskdataReportRequest{
    return &AlibabaSecurityJaqWsgriskdataReportRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqWsgriskdataReportRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.wsgriskdata.report"
}

func (r AlibabaSecurityJaqWsgriskdataReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqWsgriskdataReportRequest) SetWua(wua string) error {
    r.wua = wua
    r.Set("wua", wua)
    return nil
}

func (r AlibabaSecurityJaqWsgriskdataReportRequest) GetWua() string {
    return r.wua
}

func (r *AlibabaSecurityJaqWsgriskdataReportRequest) SetExtParam(extParam string) error {
    r.extParam = extParam
    r.Set("ext_param", extParam)
    return nil
}

func (r AlibabaSecurityJaqWsgriskdataReportRequest) GetExtParam() string {
    return r.extParam
}

