package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
应用风险扫描批量提交接口 APIRequest
alibaba.security.jaq.app.risk.scanbatch

批量提交应用进行风险扫描(含漏洞扫描、恶意代码检测),扫描完成后可通过对应的查询接口查询扫描结果
*/
type AlibabaSecurityJaqAppRiskScanbatchRequest struct {
    model.Params

    // APP信息
    appInfo   *AppInfoBatch 

    // 扫描类型
    scanTypes   []string 

}

func NewAlibabaSecurityJaqAppRiskScanbatchRequest() *AlibabaSecurityJaqAppRiskScanbatchRequest{
    return &AlibabaSecurityJaqAppRiskScanbatchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqAppRiskScanbatchRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.app.risk.scanbatch"
}

func (r AlibabaSecurityJaqAppRiskScanbatchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqAppRiskScanbatchRequest) SetAppInfo(appInfo *AppInfoBatch) error {
    r.appInfo = appInfo
    r.Set("app_info", appInfo)
    return nil
}

func (r AlibabaSecurityJaqAppRiskScanbatchRequest) GetAppInfo() *AppInfoBatch {
    return r.appInfo
}

func (r *AlibabaSecurityJaqAppRiskScanbatchRequest) SetScanTypes(scanTypes []string) error {
    r.scanTypes = scanTypes
    r.Set("scan_types", scanTypes)
    return nil
}

func (r AlibabaSecurityJaqAppRiskScanbatchRequest) GetScanTypes() []string {
    return r.scanTypes
}

