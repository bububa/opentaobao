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
    _clients   []string
    // 描述长度(9~200)
    _description   string
    // 扩展信息，JSON格式。
    _extJson   string
    // 小程序icon
    _icon   string
    // 小程序名称按平台规则自动生成，该字段仅做兜底使用。
    _name   string
    // 模板id
    _templateId   string
    // 模板版本
    _templateVersion   string
    // 小程序简称【1-16】字符，可重名
    _alias   string
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
func (r *TaobaoMiniappTemplateInstantiateRequest) SetClients(_clients []string) error {
    r._clients = _clients
    r.Set("clients", _clients)
    return nil
}

// Clients Getter
func (r TaobaoMiniappTemplateInstantiateRequest) GetClients() []string {
    return r._clients
}
// Description Setter
// 描述长度(9~200)
func (r *TaobaoMiniappTemplateInstantiateRequest) SetDescription(_description string) error {
    r._description = _description
    r.Set("description", _description)
    return nil
}

// Description Getter
func (r TaobaoMiniappTemplateInstantiateRequest) GetDescription() string {
    return r._description
}
// ExtJson Setter
// 扩展信息，JSON格式。
func (r *TaobaoMiniappTemplateInstantiateRequest) SetExtJson(_extJson string) error {
    r._extJson = _extJson
    r.Set("ext_json", _extJson)
    return nil
}

// ExtJson Getter
func (r TaobaoMiniappTemplateInstantiateRequest) GetExtJson() string {
    return r._extJson
}
// Icon Setter
// 小程序icon
func (r *TaobaoMiniappTemplateInstantiateRequest) SetIcon(_icon string) error {
    r._icon = _icon
    r.Set("icon", _icon)
    return nil
}

// Icon Getter
func (r TaobaoMiniappTemplateInstantiateRequest) GetIcon() string {
    return r._icon
}
// Name Setter
// 小程序名称按平台规则自动生成，该字段仅做兜底使用。
func (r *TaobaoMiniappTemplateInstantiateRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoMiniappTemplateInstantiateRequest) GetName() string {
    return r._name
}
// TemplateId Setter
// 模板id
func (r *TaobaoMiniappTemplateInstantiateRequest) SetTemplateId(_templateId string) error {
    r._templateId = _templateId
    r.Set("template_id", _templateId)
    return nil
}

// TemplateId Getter
func (r TaobaoMiniappTemplateInstantiateRequest) GetTemplateId() string {
    return r._templateId
}
// TemplateVersion Setter
// 模板版本
func (r *TaobaoMiniappTemplateInstantiateRequest) SetTemplateVersion(_templateVersion string) error {
    r._templateVersion = _templateVersion
    r.Set("template_version", _templateVersion)
    return nil
}

// TemplateVersion Getter
func (r TaobaoMiniappTemplateInstantiateRequest) GetTemplateVersion() string {
    return r._templateVersion
}
// Alias Setter
// 小程序简称【1-16】字符，可重名
func (r *TaobaoMiniappTemplateInstantiateRequest) SetAlias(_alias string) error {
    r._alias = _alias
    r.Set("alias", _alias)
    return nil
}

// Alias Getter
func (r TaobaoMiniappTemplateInstantiateRequest) GetAlias() string {
    return r._alias
}
