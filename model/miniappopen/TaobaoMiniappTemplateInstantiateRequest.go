package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
构建实例化应用 API请求
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

// 初始化TaobaoMiniappTemplateInstantiateRequest对象
func NewTaobaoMiniappTemplateInstantiateRequest() *TaobaoMiniappTemplateInstantiateRequest{
    return &TaobaoMiniappTemplateInstantiateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappTemplateInstantiateRequest) GetApiMethodName() string {
    return "taobao.miniapp.template.instantiate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappTemplateInstantiateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Clients Setter
// 投放端,目前可投放： taobao(淘宝),tmall(天猫)，taobao为必填，需要模板在这些端上已经发布上线
func (r *TaobaoMiniappTemplateInstantiateRequest) SetClients(clients []string) error {
    r.clients = clients
    r.Set("clients", clients)
    return nil
}

// Clients Getter
func (r TaobaoMiniappTemplateInstantiateRequest) GetClients() []string {
    return r.clients
}
// Description Setter
// 描述长度(9~200)
func (r *TaobaoMiniappTemplateInstantiateRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

// Description Getter
func (r TaobaoMiniappTemplateInstantiateRequest) GetDescription() string {
    return r.description
}
// ExtJson Setter
// 扩展信息，JSON格式。
func (r *TaobaoMiniappTemplateInstantiateRequest) SetExtJson(extJson string) error {
    r.extJson = extJson
    r.Set("ext_json", extJson)
    return nil
}

// ExtJson Getter
func (r TaobaoMiniappTemplateInstantiateRequest) GetExtJson() string {
    return r.extJson
}
// Icon Setter
// 小程序icon
func (r *TaobaoMiniappTemplateInstantiateRequest) SetIcon(icon string) error {
    r.icon = icon
    r.Set("icon", icon)
    return nil
}

// Icon Getter
func (r TaobaoMiniappTemplateInstantiateRequest) GetIcon() string {
    return r.icon
}
// Name Setter
// 小程序名称按平台规则自动生成，该字段仅做兜底使用。
func (r *TaobaoMiniappTemplateInstantiateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoMiniappTemplateInstantiateRequest) GetName() string {
    return r.name
}
// TemplateId Setter
// 模板id
func (r *TaobaoMiniappTemplateInstantiateRequest) SetTemplateId(templateId string) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

// TemplateId Getter
func (r TaobaoMiniappTemplateInstantiateRequest) GetTemplateId() string {
    return r.templateId
}
// TemplateVersion Setter
// 模板版本
func (r *TaobaoMiniappTemplateInstantiateRequest) SetTemplateVersion(templateVersion string) error {
    r.templateVersion = templateVersion
    r.Set("template_version", templateVersion)
    return nil
}

// TemplateVersion Getter
func (r TaobaoMiniappTemplateInstantiateRequest) GetTemplateVersion() string {
    return r.templateVersion
}
// Alias Setter
// 小程序简称【1-16】字符，可重名
func (r *TaobaoMiniappTemplateInstantiateRequest) SetAlias(alias string) error {
    r.alias = alias
    r.Set("alias", alias)
    return nil
}

// Alias Getter
func (r TaobaoMiniappTemplateInstantiateRequest) GetAlias() string {
    return r.alias
}
