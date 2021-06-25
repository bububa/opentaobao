package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
恶意网址检测接口 APIRequest
alibaba.security.jaq.url.scan

url扫描接口
*/
type AlibabaSecurityJaqUrlScanRequest struct {
    model.Params

    // 扫描参数
    paramUrlScanParamList   *UrlScanParamList 

}

func NewAlibabaSecurityJaqUrlScanRequest() *AlibabaSecurityJaqUrlScanRequest{
    return &AlibabaSecurityJaqUrlScanRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqUrlScanRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.url.scan"
}

func (r AlibabaSecurityJaqUrlScanRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqUrlScanRequest) SetParamUrlScanParamList(paramUrlScanParamList *UrlScanParamList) error {
    r.paramUrlScanParamList = paramUrlScanParamList
    r.Set("param_url_scan_param_list", paramUrlScanParamList)
    return nil
}

func (r AlibabaSecurityJaqUrlScanRequest) GetParamUrlScanParamList() *UrlScanParamList {
    return r.paramUrlScanParamList
}

