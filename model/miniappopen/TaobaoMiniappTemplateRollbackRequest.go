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
type TaobaoMiniappTemplateRollbackRequest struct {
    model.Params
    // 要回滚的投放端,目前可投放： taobao,tmall
    clients   []string
    // 小程序app_id
    appId   string
    // 回到到该app_version
    appVersion   string
    // 实例化小程序对应的模板id
    templateId   string
    // 与app_version对应的模板版本
    templateVersion   string
}

// 初始化TaobaoMiniappTemplateRollbackRequest对象
func NewTaobaoMiniappTemplateRollbackRequest() *TaobaoMiniappTemplateRollbackRequest{
    return &TaobaoMiniappTemplateRollbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappTemplateRollbackRequest) GetApiMethodName() string {
    return "taobao.miniapp.template.rollback"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappTemplateRollbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Clients Setter
// 要回滚的投放端,目前可投放： taobao,tmall
func (r *TaobaoMiniappTemplateRollbackRequest) SetClients(clients []string) error {
    r.clients = clients
    r.Set("clients", clients)
    return nil
}

// Clients Getter
func (r TaobaoMiniappTemplateRollbackRequest) GetClients() []string {
    return r.clients
}
// AppId Setter
// 小程序app_id
func (r *TaobaoMiniappTemplateRollbackRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

// AppId Getter
func (r TaobaoMiniappTemplateRollbackRequest) GetAppId() string {
    return r.appId
}
// AppVersion Setter
// 回到到该app_version
func (r *TaobaoMiniappTemplateRollbackRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

// AppVersion Getter
func (r TaobaoMiniappTemplateRollbackRequest) GetAppVersion() string {
    return r.appVersion
}
// TemplateId Setter
// 实例化小程序对应的模板id
func (r *TaobaoMiniappTemplateRollbackRequest) SetTemplateId(templateId string) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

// TemplateId Getter
func (r TaobaoMiniappTemplateRollbackRequest) GetTemplateId() string {
    return r.templateId
}
// TemplateVersion Setter
// 与app_version对应的模板版本
func (r *TaobaoMiniappTemplateRollbackRequest) SetTemplateVersion(templateVersion string) error {
    r.templateVersion = templateVersion
    r.Set("template_version", templateVersion)
    return nil
}

// TemplateVersion Getter
func (r TaobaoMiniappTemplateRollbackRequest) GetTemplateVersion() string {
    return r.templateVersion
}
