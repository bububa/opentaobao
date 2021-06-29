package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
下线实例化应用 API请求
taobao.miniapp.template.offlineapp

对指定的实例化小程序进行下线,需要指定clients和app_version
*/
type TaobaoMiniappTemplateOfflineappRequest struct {
    model.Params
    // 要下线的投放端,目前可投放： taobao(淘宝),tmall(天猫)
    clients   []string
    // 要下线的小程序app_id
    appId   string
    // 要下线的小程序版本号
    appVersion   string
    // 模板id
    templateId   string
}

// 初始化TaobaoMiniappTemplateOfflineappRequest对象
func NewTaobaoMiniappTemplateOfflineappRequest() *TaobaoMiniappTemplateOfflineappRequest{
    return &TaobaoMiniappTemplateOfflineappRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappTemplateOfflineappRequest) GetApiMethodName() string {
    return "taobao.miniapp.template.offlineapp"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappTemplateOfflineappRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Clients Setter
// 要下线的投放端,目前可投放： taobao(淘宝),tmall(天猫)
func (r *TaobaoMiniappTemplateOfflineappRequest) SetClients(clients []string) error {
    r.clients = clients
    r.Set("clients", clients)
    return nil
}

// Clients Getter
func (r TaobaoMiniappTemplateOfflineappRequest) GetClients() []string {
    return r.clients
}
// AppId Setter
// 要下线的小程序app_id
func (r *TaobaoMiniappTemplateOfflineappRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

// AppId Getter
func (r TaobaoMiniappTemplateOfflineappRequest) GetAppId() string {
    return r.appId
}
// AppVersion Setter
// 要下线的小程序版本号
func (r *TaobaoMiniappTemplateOfflineappRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

// AppVersion Getter
func (r TaobaoMiniappTemplateOfflineappRequest) GetAppVersion() string {
    return r.appVersion
}
// TemplateId Setter
// 模板id
func (r *TaobaoMiniappTemplateOfflineappRequest) SetTemplateId(templateId string) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

// TemplateId Getter
func (r TaobaoMiniappTemplateOfflineappRequest) GetTemplateId() string {
    return r.templateId
}
