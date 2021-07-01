package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（已废弃）更新实例化应用 API请求
taobao.miniapp.template.update

商家应用c端模板实例化小程序更新
*/
type TaobaoMiniappTemplateUpdateAPIRequest struct {
    model.Params
    // 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
    _clients   []string
    // 应用id
    _id   string
    // schema信息，不填且 应用线上版本使用的templateId与传入的templateId不一致，则会报错; 一致，则复用线上版本的schema。
    _extJson   string
    // 模板id
    _templateId   string
    // 模板版本
    _templateVersion   string
}

// 初始化TaobaoMiniappTemplateUpdateAPIRequest对象
func NewTaobaoMiniappTemplateUpdateRequest() *TaobaoMiniappTemplateUpdateAPIRequest{
    return &TaobaoMiniappTemplateUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappTemplateUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.miniapp.template.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappTemplateUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Clients Setter
// 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
func (r *TaobaoMiniappTemplateUpdateAPIRequest) SetClients(_clients []string) error {
    r._clients = _clients
    r.Set("clients", _clients)
    return nil
}

// Clients Getter
func (r TaobaoMiniappTemplateUpdateAPIRequest) GetClients() []string {
    return r._clients
}
// Id Setter
// 应用id
func (r *TaobaoMiniappTemplateUpdateAPIRequest) SetId(_id string) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r TaobaoMiniappTemplateUpdateAPIRequest) GetId() string {
    return r._id
}
// ExtJson Setter
// schema信息，不填且 应用线上版本使用的templateId与传入的templateId不一致，则会报错; 一致，则复用线上版本的schema。
func (r *TaobaoMiniappTemplateUpdateAPIRequest) SetExtJson(_extJson string) error {
    r._extJson = _extJson
    r.Set("ext_json", _extJson)
    return nil
}

// ExtJson Getter
func (r TaobaoMiniappTemplateUpdateAPIRequest) GetExtJson() string {
    return r._extJson
}
// TemplateId Setter
// 模板id
func (r *TaobaoMiniappTemplateUpdateAPIRequest) SetTemplateId(_templateId string) error {
    r._templateId = _templateId
    r.Set("template_id", _templateId)
    return nil
}

// TemplateId Getter
func (r TaobaoMiniappTemplateUpdateAPIRequest) GetTemplateId() string {
    return r._templateId
}
// TemplateVersion Setter
// 模板版本
func (r *TaobaoMiniappTemplateUpdateAPIRequest) SetTemplateVersion(_templateVersion string) error {
    r._templateVersion = _templateVersion
    r.Set("template_version", _templateVersion)
    return nil
}

// TemplateVersion Getter
func (r TaobaoMiniappTemplateUpdateAPIRequest) GetTemplateVersion() string {
    return r._templateVersion
}
