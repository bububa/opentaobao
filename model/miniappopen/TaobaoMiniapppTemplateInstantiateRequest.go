package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（已废弃）构建实例化应用 API请求
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

// 初始化TaobaoMiniapppTemplateInstantiateRequest对象
func NewTaobaoMiniapppTemplateInstantiateRequest() *TaobaoMiniapppTemplateInstantiateRequest{
    return &TaobaoMiniapppTemplateInstantiateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniapppTemplateInstantiateRequest) GetApiMethodName() string {
    return "taobao.miniappp.template.instantiate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniapppTemplateInstantiateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Clients Setter
// 投放端,目前可投放： taobao(淘宝),tmall(天猫)
func (r *TaobaoMiniapppTemplateInstantiateRequest) SetClients(clients []string) error {
    r.clients = clients
    r.Set("clients", clients)
    return nil
}

// Clients Getter
func (r TaobaoMiniapppTemplateInstantiateRequest) GetClients() []string {
    return r.clients
}
// Description Setter
// 描述长度(9~200)
func (r *TaobaoMiniapppTemplateInstantiateRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

// Description Getter
func (r TaobaoMiniapppTemplateInstantiateRequest) GetDescription() string {
    return r.description
}
// ExtJson Setter
// schemadata, json字符串
func (r *TaobaoMiniapppTemplateInstantiateRequest) SetExtJson(extJson string) error {
    r.extJson = extJson
    r.Set("ext_json", extJson)
    return nil
}

// ExtJson Getter
func (r TaobaoMiniapppTemplateInstantiateRequest) GetExtJson() string {
    return r.extJson
}
// Icon Setter
// 小程序icon
func (r *TaobaoMiniapppTemplateInstantiateRequest) SetIcon(icon string) error {
    r.icon = icon
    r.Set("icon", icon)
    return nil
}

// Icon Getter
func (r TaobaoMiniapppTemplateInstantiateRequest) GetIcon() string {
    return r.icon
}
// Name Setter
// 小程序名称
func (r *TaobaoMiniapppTemplateInstantiateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoMiniapppTemplateInstantiateRequest) GetName() string {
    return r.name
}
// TemplateId Setter
// 模板id
func (r *TaobaoMiniapppTemplateInstantiateRequest) SetTemplateId(templateId string) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

// TemplateId Getter
func (r TaobaoMiniapppTemplateInstantiateRequest) GetTemplateId() string {
    return r.templateId
}
// TemplateVersion Setter
// 模板版本
func (r *TaobaoMiniapppTemplateInstantiateRequest) SetTemplateVersion(templateVersion string) error {
    r.templateVersion = templateVersion
    r.Set("template_version", templateVersion)
    return nil
}

// TemplateVersion Getter
func (r TaobaoMiniapppTemplateInstantiateRequest) GetTemplateVersion() string {
    return r.templateVersion
}
