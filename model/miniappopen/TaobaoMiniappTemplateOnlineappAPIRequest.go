package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上线实例化应用 API请求
taobao.miniapp.template.onlineapp

将指定的预览版本发布上线，预览版本号由构建实例化或更新实例化接口返回。
*/
type TaobaoMiniappTemplateOnlineappAPIRequest struct {
    model.Params
    // 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
    _clients   []string
    // 小程序app_id
    _appId   string
    // 模板id
    _templateId   string
    // 模板版本
    _templateVersion   string
    // 待上线的预览版本号
    _appVersion   string
}

// 初始化TaobaoMiniappTemplateOnlineappAPIRequest对象
func NewTaobaoMiniappTemplateOnlineappRequest() *TaobaoMiniappTemplateOnlineappAPIRequest{
    return &TaobaoMiniappTemplateOnlineappAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappTemplateOnlineappAPIRequest) GetApiMethodName() string {
    return "taobao.miniapp.template.onlineapp"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappTemplateOnlineappAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Clients Setter
// 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
func (r *TaobaoMiniappTemplateOnlineappAPIRequest) SetClients(_clients []string) error {
    r._clients = _clients
    r.Set("clients", _clients)
    return nil
}

// Clients Getter
func (r TaobaoMiniappTemplateOnlineappAPIRequest) GetClients() []string {
    return r._clients
}
// AppId Setter
// 小程序app_id
func (r *TaobaoMiniappTemplateOnlineappAPIRequest) SetAppId(_appId string) error {
    r._appId = _appId
    r.Set("app_id", _appId)
    return nil
}

// AppId Getter
func (r TaobaoMiniappTemplateOnlineappAPIRequest) GetAppId() string {
    return r._appId
}
// TemplateId Setter
// 模板id
func (r *TaobaoMiniappTemplateOnlineappAPIRequest) SetTemplateId(_templateId string) error {
    r._templateId = _templateId
    r.Set("template_id", _templateId)
    return nil
}

// TemplateId Getter
func (r TaobaoMiniappTemplateOnlineappAPIRequest) GetTemplateId() string {
    return r._templateId
}
// TemplateVersion Setter
// 模板版本
func (r *TaobaoMiniappTemplateOnlineappAPIRequest) SetTemplateVersion(_templateVersion string) error {
    r._templateVersion = _templateVersion
    r.Set("template_version", _templateVersion)
    return nil
}

// TemplateVersion Getter
func (r TaobaoMiniappTemplateOnlineappAPIRequest) GetTemplateVersion() string {
    return r._templateVersion
}
// AppVersion Setter
// 待上线的预览版本号
func (r *TaobaoMiniappTemplateOnlineappAPIRequest) SetAppVersion(_appVersion string) error {
    r._appVersion = _appVersion
    r.Set("app_version", _appVersion)
    return nil
}

// AppVersion Getter
func (r TaobaoMiniappTemplateOnlineappAPIRequest) GetAppVersion() string {
    return r._appVersion
}
