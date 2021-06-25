package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
应用风险扫描提交接口 APIRequest
alibaba.security.jaq.app.risk.scan

提交应用进行风险扫描(含漏洞扫描、恶意代码检测、仿冒监测),扫描完成后可通过对应的查询接口查询扫描结果
*/
type AlibabaSecurityJaqAppRiskScanRequest struct {
    model.Params

    // 应用信息
    appInfo   *ScanAppInfo 

    // 扫描类型：vuln-漏洞扫描 malware-恶意代码检测 fake-仿冒监测 plugin-插件扫描 注: dataType为2时 不支持 仿冒监测
    scanTypes   []String 

    // 额外的信息，根据具体业务定
    extParam   string 

}

func NewAlibabaSecurityJaqAppRiskScanRequest() *AlibabaSecurityJaqAppRiskScanRequest{
    return &AlibabaSecurityJaqAppRiskScanRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqAppRiskScanRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.app.risk.scan"
}

func (r AlibabaSecurityJaqAppRiskScanRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqAppRiskScanRequest) SetAppInfo(appInfo *ScanAppInfo) error {
    r.appInfo = appInfo
    r.Set("app_info", appInfo)
    return nil
}

func (r AlibabaSecurityJaqAppRiskScanRequest) GetAppInfo() *ScanAppInfo {
    return r.appInfo
}

func (r *AlibabaSecurityJaqAppRiskScanRequest) SetScanTypes(scanTypes []String) error {
    r.scanTypes = scanTypes
    r.Set("scan_types", scanTypes)
    return nil
}

func (r AlibabaSecurityJaqAppRiskScanRequest) GetScanTypes() []String {
    return r.scanTypes
}

func (r *AlibabaSecurityJaqAppRiskScanRequest) SetExtParam(extParam string) error {
    r.extParam = extParam
    r.Set("ext_param", extParam)
    return nil
}

func (r AlibabaSecurityJaqAppRiskScanRequest) GetExtParam() string {
    return r.extParam
}

