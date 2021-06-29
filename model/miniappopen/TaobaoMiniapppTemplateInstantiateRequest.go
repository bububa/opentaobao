package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（已废弃）构建实例化应用 APIRequest
taobao.miniappp.template.instantiate

实例化saas化的小程序
*/
type TaobaoMiniapppTemplateInstantiateRequest struct {
    model.Params

    // 投放端,目前可投放： taobao(淘宝),tmall(天猫)
    clients   []string 

    // 描述长度(9~200)
    description   string 

    // schemadata, json字符串
    extJson   string 

    // 小程序icon
    icon   string 

    // 小程序名称
    name   string 

    // 模板id
    templateId   string 

    // 模板版本
    templateVersion   string 

}

func NewTaobaoMiniapppTemplateInstantiateRequest() *TaobaoMiniapppTemplateInstantiateRequest{
    return &TaobaoMiniapppTemplateInstantiateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniapppTemplateInstantiateRequest) GetApiMethodName() string {
    return "taobao.miniappp.template.instantiate"
}

func (r TaobaoMiniapppTemplateInstantiateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniapppTemplateInstantiateRequest) SetClients(clients []string) error {
    r.clients = clients
    r.Set("clients", clients)
    return nil
}

func (r TaobaoMiniapppTemplateInstantiateRequest) GetClients() []string {
    return r.clients
}

func (r *TaobaoMiniapppTemplateInstantiateRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

func (r TaobaoMiniapppTemplateInstantiateRequest) GetDescription() string {
    return r.description
}

func (r *TaobaoMiniapppTemplateInstantiateRequest) SetExtJson(extJson string) error {
    r.extJson = extJson
    r.Set("ext_json", extJson)
    return nil
}

func (r TaobaoMiniapppTemplateInstantiateRequest) GetExtJson() string {
    return r.extJson
}

func (r *TaobaoMiniapppTemplateInstantiateRequest) SetIcon(icon string) error {
    r.icon = icon
    r.Set("icon", icon)
    return nil
}

func (r TaobaoMiniapppTemplateInstantiateRequest) GetIcon() string {
    return r.icon
}

func (r *TaobaoMiniapppTemplateInstantiateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoMiniapppTemplateInstantiateRequest) GetName() string {
    return r.name
}

func (r *TaobaoMiniapppTemplateInstantiateRequest) SetTemplateId(templateId string) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

func (r TaobaoMiniapppTemplateInstantiateRequest) GetTemplateId() string {
    return r.templateId
}

func (r *TaobaoMiniapppTemplateInstantiateRequest) SetTemplateVersion(templateVersion string) error {
    r.templateVersion = templateVersion
    r.Set("template_version", templateVersion)
    return nil
}

func (r TaobaoMiniapppTemplateInstantiateRequest) GetTemplateVersion() string {
    return r.templateVersion
}

