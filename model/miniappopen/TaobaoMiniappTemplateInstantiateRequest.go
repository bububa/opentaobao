package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
构建实例化应用 APIRequest
taobao.miniapp.template.instantiate

实例化saas化的小程序
*/
type TaobaoMiniappTemplateInstantiateRequest struct {
    model.Params

    // 投放端,目前可投放： taobao(淘宝),tmall(天猫)，taobao为必填，需要模板在这些端上已经发布上线
    clients   []string 

    // 描述长度(9~200)
    description   string 

    // 扩展信息，JSON格式。
    extJson   string 

    // 小程序icon
    icon   string 

    // 小程序名称按平台规则自动生成，该字段仅做兜底使用。
    name   string 

    // 模板id
    templateId   string 

    // 模板版本
    templateVersion   string 

    // 小程序简称【1-16】字符，可重名
    alias   string 

}

func NewTaobaoMiniappTemplateInstantiateRequest() *TaobaoMiniappTemplateInstantiateRequest{
    return &TaobaoMiniappTemplateInstantiateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappTemplateInstantiateRequest) GetApiMethodName() string {
    return "taobao.miniapp.template.instantiate"
}

func (r TaobaoMiniappTemplateInstantiateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappTemplateInstantiateRequest) SetClients(clients []string) error {
    r.clients = clients
    r.Set("clients", clients)
    return nil
}

func (r TaobaoMiniappTemplateInstantiateRequest) GetClients() []string {
    return r.clients
}

func (r *TaobaoMiniappTemplateInstantiateRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

func (r TaobaoMiniappTemplateInstantiateRequest) GetDescription() string {
    return r.description
}

func (r *TaobaoMiniappTemplateInstantiateRequest) SetExtJson(extJson string) error {
    r.extJson = extJson
    r.Set("ext_json", extJson)
    return nil
}

func (r TaobaoMiniappTemplateInstantiateRequest) GetExtJson() string {
    return r.extJson
}

func (r *TaobaoMiniappTemplateInstantiateRequest) SetIcon(icon string) error {
    r.icon = icon
    r.Set("icon", icon)
    return nil
}

func (r TaobaoMiniappTemplateInstantiateRequest) GetIcon() string {
    return r.icon
}

func (r *TaobaoMiniappTemplateInstantiateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoMiniappTemplateInstantiateRequest) GetName() string {
    return r.name
}

func (r *TaobaoMiniappTemplateInstantiateRequest) SetTemplateId(templateId string) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

func (r TaobaoMiniappTemplateInstantiateRequest) GetTemplateId() string {
    return r.templateId
}

func (r *TaobaoMiniappTemplateInstantiateRequest) SetTemplateVersion(templateVersion string) error {
    r.templateVersion = templateVersion
    r.Set("template_version", templateVersion)
    return nil
}

func (r TaobaoMiniappTemplateInstantiateRequest) GetTemplateVersion() string {
    return r.templateVersion
}

func (r *TaobaoMiniappTemplateInstantiateRequest) SetAlias(alias string) error {
    r.alias = alias
    r.Set("alias", alias)
    return nil
}

func (r TaobaoMiniappTemplateInstantiateRequest) GetAlias() string {
    return r.alias
}

