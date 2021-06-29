package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新实例化应用 API请求
taobao.miniapp.template.updateapp

商家应用c端模板实例化小程序更新，生成新的版本，但不会自动上线新版本
*/
type TaobaoMiniappTemplateUpdateappRequest struct {
    model.Params
    // 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
    _clients   []string
    // 应用id，如果应用在对应端上已有的最新版本所使用的模板id,模板version和extjson，与本次入参一致，则认为不需要更新，返回已有的版本。
    _appId   string
    // 扩展信息。线上版本使用的template_id与传入的template_id一致时，可不填。
    _extJson   string
    // 模板id
    _templateId   string
    // 模板版本
    _templateVersion   string
    // Logo更新，1月5次
    _icon   string
    // 描述更新，1年5次
    _desc   string
    // 简称更新，1年5次
    _alias   string
}

// 初始化TaobaoMiniappTemplateUpdateappRequest对象
func NewTaobaoMiniappTemplateUpdateappRequest() *TaobaoMiniappTemplateUpdateappRequest{
    return &TaobaoMiniappTemplateUpdateappRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappTemplateUpdateappRequest) GetApiMethodName() string {
    return "taobao.miniapp.template.updateapp"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappTemplateUpdateappRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Clients Setter
// 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
func (r *TaobaoMiniappTemplateUpdateappRequest) SetClients(_clients []string) error {
    r._clients = _clients
    r.Set("clients", _clients)
    return nil
}

// Clients Getter
func (r TaobaoMiniappTemplateUpdateappRequest) GetClients() []string {
    return r._clients
}
// AppId Setter
// 应用id，如果应用在对应端上已有的最新版本所使用的模板id,模板version和extjson，与本次入参一致，则认为不需要更新，返回已有的版本。
func (r *TaobaoMiniappTemplateUpdateappRequest) SetAppId(_appId string) error {
    r._appId = _appId
    r.Set("app_id", _appId)
    return nil
}

// AppId Getter
func (r TaobaoMiniappTemplateUpdateappRequest) GetAppId() string {
    return r._appId
}
// ExtJson Setter
// 扩展信息。线上版本使用的template_id与传入的template_id一致时，可不填。
func (r *TaobaoMiniappTemplateUpdateappRequest) SetExtJson(_extJson string) error {
    r._extJson = _extJson
    r.Set("ext_json", _extJson)
    return nil
}

// ExtJson Getter
func (r TaobaoMiniappTemplateUpdateappRequest) GetExtJson() string {
    return r._extJson
}
// TemplateId Setter
// 模板id
func (r *TaobaoMiniappTemplateUpdateappRequest) SetTemplateId(_templateId string) error {
    r._templateId = _templateId
    r.Set("template_id", _templateId)
    return nil
}

// TemplateId Getter
func (r TaobaoMiniappTemplateUpdateappRequest) GetTemplateId() string {
    return r._templateId
}
// TemplateVersion Setter
// 模板版本
func (r *TaobaoMiniappTemplateUpdateappRequest) SetTemplateVersion(_templateVersion string) error {
    r._templateVersion = _templateVersion
    r.Set("template_version", _templateVersion)
    return nil
}

// TemplateVersion Getter
func (r TaobaoMiniappTemplateUpdateappRequest) GetTemplateVersion() string {
    return r._templateVersion
}
// Icon Setter
// Logo更新，1月5次
func (r *TaobaoMiniappTemplateUpdateappRequest) SetIcon(_icon string) error {
    r._icon = _icon
    r.Set("icon", _icon)
    return nil
}

// Icon Getter
func (r TaobaoMiniappTemplateUpdateappRequest) GetIcon() string {
    return r._icon
}
// Desc Setter
// 描述更新，1年5次
func (r *TaobaoMiniappTemplateUpdateappRequest) SetDesc(_desc string) error {
    r._desc = _desc
    r.Set("desc", _desc)
    return nil
}

// Desc Getter
func (r TaobaoMiniappTemplateUpdateappRequest) GetDesc() string {
    return r._desc
}
// Alias Setter
// 简称更新，1年5次
func (r *TaobaoMiniappTemplateUpdateappRequest) SetAlias(_alias string) error {
    r._alias = _alias
    r.Set("alias", _alias)
    return nil
}

// Alias Getter
func (r TaobaoMiniappTemplateUpdateappRequest) GetAlias() string {
    return r._alias
}
