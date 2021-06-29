package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
恶意网址检测接口 API请求
alibaba.security.jaq.url.scan

url扫描接口
*/
type AlibabaSecurityJaqUrlScanRequest struct {
    model.Params
    // 扫描参数
    paramUrlScanParamList   *UrlScanParamList
}

// 初始化AlibabaSecurityJaqUrlScanRequest对象
func NewAlibabaSecurityJaqUrlScanRequest() *AlibabaSecurityJaqUrlScanRequest{
    return &AlibabaSecurityJaqUrlScanRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqUrlScanRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.url.scan"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqUrlScanRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamUrlScanParamList Setter
// 扫描参数
func (r *AlibabaSecurityJaqUrlScanRequest) SetParamUrlScanParamList(paramUrlScanParamList *UrlScanParamList) error {
    r.paramUrlScanParamList = paramUrlScanParamList
    r.Set("param_url_scan_param_list", paramUrlScanParamList)
    return nil
}

// ParamUrlScanParamList Getter
func (r AlibabaSecurityJaqUrlScanRequest) GetParamUrlScanParamList() *UrlScanParamList {
    return r.paramUrlScanParamList
}
