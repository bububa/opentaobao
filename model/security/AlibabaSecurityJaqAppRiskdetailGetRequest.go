package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
应用风险详细信息查询接口 APIRequest
alibaba.security.jaq.app.riskdetail.get

用户通过alibaba.security.jaq.app.risk.scan接口提交应用进行风险扫描后，用此接口获取风险详细信息,包含漏洞列表、恶意代码列表、仿冒应用列表等信息
*/
type AlibabaSecurityJaqAppRiskdetailGetRequest struct {
    model.Params

    // 任务唯一标识
    itemId   string 

    // 本地化语言信息
    locale   *Locale 

}

func NewAlibabaSecurityJaqAppRiskdetailGetRequest() *AlibabaSecurityJaqAppRiskdetailGetRequest{
    return &AlibabaSecurityJaqAppRiskdetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqAppRiskdetailGetRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.app.riskdetail.get"
}

func (r AlibabaSecurityJaqAppRiskdetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqAppRiskdetailGetRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlibabaSecurityJaqAppRiskdetailGetRequest) GetItemId() string {
    return r.itemId
}

func (r *AlibabaSecurityJaqAppRiskdetailGetRequest) SetLocale(locale *Locale) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

func (r AlibabaSecurityJaqAppRiskdetailGetRequest) GetLocale() *Locale {
    return r.locale
}

