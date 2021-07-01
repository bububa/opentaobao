package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
应用风险扫描批量提交接口 API请求
alibaba.security.jaq.app.risk.scanbatch

批量提交应用进行风险扫描(含漏洞扫描、恶意代码检测),扫描完成后可通过对应的查询接口查询扫描结果
*/
type AlibabaSecurityJaqAppRiskScanbatchAPIRequest struct {
    model.Params
    // APP信息
    _appInfo   *AppInfoBatch
    // 扫描类型
    _scanTypes   []string
}

// 初始化AlibabaSecurityJaqAppRiskScanbatchAPIRequest对象
func NewAlibabaSecurityJaqAppRiskScanbatchRequest() *AlibabaSecurityJaqAppRiskScanbatchAPIRequest{
    return &AlibabaSecurityJaqAppRiskScanbatchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqAppRiskScanbatchAPIRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.app.risk.scanbatch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqAppRiskScanbatchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppInfo Setter
// APP信息
func (r *AlibabaSecurityJaqAppRiskScanbatchAPIRequest) SetAppInfo(_appInfo *AppInfoBatch) error {
    r._appInfo = _appInfo
    r.Set("app_info", _appInfo)
    return nil
}

// AppInfo Getter
func (r AlibabaSecurityJaqAppRiskScanbatchAPIRequest) GetAppInfo() *AppInfoBatch {
    return r._appInfo
}
// ScanTypes Setter
// 扫描类型
func (r *AlibabaSecurityJaqAppRiskScanbatchAPIRequest) SetScanTypes(_scanTypes []string) error {
    r._scanTypes = _scanTypes
    r.Set("scan_types", _scanTypes)
    return nil
}

// ScanTypes Getter
func (r AlibabaSecurityJaqAppRiskScanbatchAPIRequest) GetScanTypes() []string {
    return r._scanTypes
}
