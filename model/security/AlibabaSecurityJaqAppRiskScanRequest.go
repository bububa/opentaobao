package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
应用风险扫描提交接口 API请求
alibaba.security.jaq.app.risk.scan

提交应用进行风险扫描(含漏洞扫描、恶意代码检测、仿冒监测),扫描完成后可通过对应的查询接口查询扫描结果
*/
type AlibabaSecurityJaqAppRiskScanRequest struct {
    model.Params
    // 应用信息
    appInfo   *ScanAppInfo
    // 扫描类型：vuln-漏洞扫描 malware-恶意代码检测 fake-仿冒监测 plugin-插件扫描 注: dataType为2时 不支持 仿冒监测
    scanTypes   []string
    // 额外的信息，根据具体业务定
    extParam   string
}

// 初始化AlibabaSecurityJaqAppRiskScanRequest对象
func NewAlibabaSecurityJaqAppRiskScanRequest() *AlibabaSecurityJaqAppRiskScanRequest{
    return &AlibabaSecurityJaqAppRiskScanRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqAppRiskScanRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.app.risk.scan"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqAppRiskScanRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppInfo Setter
// 应用信息
func (r *AlibabaSecurityJaqAppRiskScanRequest) SetAppInfo(appInfo *ScanAppInfo) error {
    r.appInfo = appInfo
    r.Set("app_info", appInfo)
    return nil
}

// AppInfo Getter
func (r AlibabaSecurityJaqAppRiskScanRequest) GetAppInfo() *ScanAppInfo {
    return r.appInfo
}
// ScanTypes Setter
// 扫描类型：vuln-漏洞扫描 malware-恶意代码检测 fake-仿冒监测 plugin-插件扫描 注: dataType为2时 不支持 仿冒监测
func (r *AlibabaSecurityJaqAppRiskScanRequest) SetScanTypes(scanTypes []string) error {
    r.scanTypes = scanTypes
    r.Set("scan_types", scanTypes)
    return nil
}

// ScanTypes Getter
func (r AlibabaSecurityJaqAppRiskScanRequest) GetScanTypes() []string {
    return r.scanTypes
}
// ExtParam Setter
// 额外的信息，根据具体业务定
func (r *AlibabaSecurityJaqAppRiskScanRequest) SetExtParam(extParam string) error {
    r.extParam = extParam
    r.Set("ext_param", extParam)
    return nil
}

// ExtParam Getter
func (r AlibabaSecurityJaqAppRiskScanRequest) GetExtParam() string {
    return r.extParam
}
