package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回滚实例化应用 API请求
taobao.miniapp.template.rollback

将实例化小程序回滚到指定版本
*/
type TaobaoMiniappTemplateRollbackAPIRequest struct {
    model.Params
    // 要回滚的投放端,目前可投放： taobao,tmall
    _clients   []string
    // 小程序app_id
    _appId   string
    // 回到到该app_version
    _appVersion   string
    // 实例化小程序对应的模板id
    _templateId   string
    // 与app_version对应的模板版本
    _templateVersion   string
}

// 初始化TaobaoMiniappTemplateRollbackAPIRequest对象
func NewTaobaoMiniappTemplateRollbackRequest() *TaobaoMiniappTemplateRollbackAPIRequest{
    return &TaobaoMiniappTemplateRollbackAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappTemplateRollbackAPIRequest) GetApiMethodName() string {
    return "taobao.miniapp.template.rollback"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappTemplateRollbackAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Clients Setter
// 要回滚的投放端,目前可投放： taobao,tmall
func (r *TaobaoMiniappTemplateRollbackAPIRequest) SetClients(_clients []string) error {
    r._clients = _clients
    r.Set("clients", _clients)
    return nil
}

// Clients Getter
func (r TaobaoMiniappTemplateRollbackAPIRequest) GetClients() []string {
    return r._clients
}
// AppId Setter
// 小程序app_id
func (r *TaobaoMiniappTemplateRollbackAPIRequest) SetAppId(_appId string) error {
    r._appId = _appId
    r.Set("app_id", _appId)
    return nil
}

// AppId Getter
func (r TaobaoMiniappTemplateRollbackAPIRequest) GetAppId() string {
    return r._appId
}
// AppVersion Setter
// 回到到该app_version
func (r *TaobaoMiniappTemplateRollbackAPIRequest) SetAppVersion(_appVersion string) error {
    r._appVersion = _appVersion
    r.Set("app_version", _appVersion)
    return nil
}

// AppVersion Getter
func (r TaobaoMiniappTemplateRollbackAPIRequest) GetAppVersion() string {
    return r._appVersion
}
// TemplateId Setter
// 实例化小程序对应的模板id
func (r *TaobaoMiniappTemplateRollbackAPIRequest) SetTemplateId(_templateId string) error {
    r._templateId = _templateId
    r.Set("template_id", _templateId)
    return nil
}

// TemplateId Getter
func (r TaobaoMiniappTemplateRollbackAPIRequest) GetTemplateId() string {
    return r._templateId
}
// TemplateVersion Setter
// 与app_version对应的模板版本
func (r *TaobaoMiniappTemplateRollbackAPIRequest) SetTemplateVersion(_templateVersion string) error {
    r._templateVersion = _templateVersion
    r.Set("template_version", _templateVersion)
    return nil
}

// TemplateVersion Getter
func (r TaobaoMiniappTemplateRollbackAPIRequest) GetTemplateVersion() string {
    return r._templateVersion
}
