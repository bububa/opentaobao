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
    _clients   []string
    // 描述长度(9~200)
    _description   string
    // schemadata, json字符串
    _extJson   string
    // 小程序icon
    _icon   string
    // 小程序名称
    _name   string
    // 模板id
    _templateId   string
    // 模板版本
    _templateVersion   string
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
func (r *TaobaoMiniapppTemplateInstantiateRequest) SetClients(_clients []string) error {
    r._clients = _clients
    r.Set("clients", _clients)
    return nil
}

// Clients Getter
func (r TaobaoMiniapppTemplateInstantiateRequest) GetClients() []string {
    return r._clients
}
// Description Setter
// 描述长度(9~200)
func (r *TaobaoMiniapppTemplateInstantiateRequest) SetDescription(_description string) error {
    r._description = _description
    r.Set("description", _description)
    return nil
}

// Description Getter
func (r TaobaoMiniapppTemplateInstantiateRequest) GetDescription() string {
    return r._description
}
// ExtJson Setter
// schemadata, json字符串
func (r *TaobaoMiniapppTemplateInstantiateRequest) SetExtJson(_extJson string) error {
    r._extJson = _extJson
    r.Set("ext_json", _extJson)
    return nil
}

// ExtJson Getter
func (r TaobaoMiniapppTemplateInstantiateRequest) GetExtJson() string {
    return r._extJson
}
// Icon Setter
// 小程序icon
func (r *TaobaoMiniapppTemplateInstantiateRequest) SetIcon(_icon string) error {
    r._icon = _icon
    r.Set("icon", _icon)
    return nil
}

// Icon Getter
func (r TaobaoMiniapppTemplateInstantiateRequest) GetIcon() string {
    return r._icon
}
// Name Setter
// 小程序名称
func (r *TaobaoMiniapppTemplateInstantiateRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoMiniapppTemplateInstantiateRequest) GetName() string {
    return r._name
}
// TemplateId Setter
// 模板id
func (r *TaobaoMiniapppTemplateInstantiateRequest) SetTemplateId(_templateId string) error {
    r._templateId = _templateId
    r.Set("template_id", _templateId)
    return nil
}

// TemplateId Getter
func (r TaobaoMiniapppTemplateInstantiateRequest) GetTemplateId() string {
    return r._templateId
}
// TemplateVersion Setter
// 模板版本
func (r *TaobaoMiniapppTemplateInstantiateRequest) SetTemplateVersion(_templateVersion string) error {
    r._templateVersion = _templateVersion
    r.Set("template_version", _templateVersion)
    return nil
}

// TemplateVersion Getter
func (r TaobaoMiniapppTemplateInstantiateRequest) GetTemplateVersion() string {
    return r._templateVersion
}
