package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
应用风险详细信息查询接口 API请求
alibaba.security.jaq.app.riskdetail.get

用户通过alibaba.security.jaq.app.risk.scan接口提交应用进行风险扫描后，用此接口获取风险详细信息,包含漏洞列表、恶意代码列表、仿冒应用列表等信息
*/
type AlibabaSecurityJaqAppRiskdetailGetRequest struct {
    model.Params
    // 任务唯一标识
    _itemId   string
    // 本地化语言信息
    _locale   *Locale
}

// 初始化AlibabaSecurityJaqAppRiskdetailGetRequest对象
func NewAlibabaSecurityJaqAppRiskdetailGetRequest() *AlibabaSecurityJaqAppRiskdetailGetRequest{
    return &AlibabaSecurityJaqAppRiskdetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqAppRiskdetailGetRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.app.riskdetail.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqAppRiskdetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 任务唯一标识
func (r *AlibabaSecurityJaqAppRiskdetailGetRequest) SetItemId(_itemId string) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r AlibabaSecurityJaqAppRiskdetailGetRequest) GetItemId() string {
    return r._itemId
}
// Locale Setter
// 本地化语言信息
func (r *AlibabaSecurityJaqAppRiskdetailGetRequest) SetLocale(_locale *Locale) error {
    r._locale = _locale
    r.Set("locale", _locale)
    return nil
}

// Locale Getter
func (r AlibabaSecurityJaqAppRiskdetailGetRequest) GetLocale() *Locale {
    return r._locale
}
