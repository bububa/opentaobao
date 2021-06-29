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
type TaobaoMiniappTemplateOnlineappRequest struct {
    model.Params
    // 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
    clients   []string
    // 小程序app_id
    appId   string
    // 模板id
    templateId   string
    // 模板版本
    templateVersion   string
    // 待上线的预览版本号
    appVersion   string
}

// 初始化TaobaoMiniappTemplateOnlineappRequest对象
func NewTaobaoMiniappTemplateOnlineappRequest() *TaobaoMiniappTemplateOnlineappRequest{
    return &TaobaoMiniappTemplateOnlineappRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappTemplateOnlineappRequest) GetApiMethodName() string {
    return "taobao.miniapp.template.onlineapp"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappTemplateOnlineappRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Clients Setter
// 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
func (r *TaobaoMiniappTemplateOnlineappRequest) SetClients(clients []string) error {
    r.clients = clients
    r.Set("clients", clients)
    return nil
}

// Clients Getter
func (r TaobaoMiniappTemplateOnlineappRequest) GetClients() []string {
    return r.clients
}
// AppId Setter
// 小程序app_id
func (r *TaobaoMiniappTemplateOnlineappRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

// AppId Getter
func (r TaobaoMiniappTemplateOnlineappRequest) GetAppId() string {
    return r.appId
}
// TemplateId Setter
// 模板id
func (r *TaobaoMiniappTemplateOnlineappRequest) SetTemplateId(templateId string) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

// TemplateId Getter
func (r TaobaoMiniappTemplateOnlineappRequest) GetTemplateId() string {
    return r.templateId
}
// TemplateVersion Setter
// 模板版本
func (r *TaobaoMiniappTemplateOnlineappRequest) SetTemplateVersion(templateVersion string) error {
    r.templateVersion = templateVersion
    r.Set("template_version", templateVersion)
    return nil
}

// TemplateVersion Getter
func (r TaobaoMiniappTemplateOnlineappRequest) GetTemplateVersion() string {
    return r.templateVersion
}
// AppVersion Setter
// 待上线的预览版本号
func (r *TaobaoMiniappTemplateOnlineappRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

// AppVersion Getter
func (r TaobaoMiniappTemplateOnlineappRequest) GetAppVersion() string {
    return r.appVersion
}
