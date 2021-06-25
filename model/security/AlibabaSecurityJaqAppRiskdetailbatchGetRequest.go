package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
应用风险详细信息批量查询接口 APIRequest
alibaba.security.jaq.app.riskdetailbatch.get

用户通过alibaba.security.jaq.app.risk.scanbatch接口提交应用进行风险批量扫描后，用此接口批量获取风险详细信息,包含漏洞列表、恶意代码列表、仿冒应用列表等信息
*/
type AlibabaSecurityJaqAppRiskdetailbatchGetRequest struct {
    model.Params

    // 任务唯一标识
    itemId   string 

    // 本地化语言信息,用于指定返回结果内容所使用的语言(默认为zh_CN,目前仅支持zh_CN)
    locale   *Locale 

}

func NewAlibabaSecurityJaqAppRiskdetailbatchGetRequest() *AlibabaSecurityJaqAppRiskdetailbatchGetRequest{
    return &AlibabaSecurityJaqAppRiskdetailbatchGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqAppRiskdetailbatchGetRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.app.riskdetailbatch.get"
}

func (r AlibabaSecurityJaqAppRiskdetailbatchGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqAppRiskdetailbatchGetRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlibabaSecurityJaqAppRiskdetailbatchGetRequest) GetItemId() string {
    return r.itemId
}

func (r *AlibabaSecurityJaqAppRiskdetailbatchGetRequest) SetLocale(locale *Locale) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

func (r AlibabaSecurityJaqAppRiskdetailbatchGetRequest) GetLocale() *Locale {
    return r.locale
}

