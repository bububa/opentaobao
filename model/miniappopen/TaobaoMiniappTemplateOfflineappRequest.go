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
    _clients   []string
    // 要下线的小程序app_id
    _appId   string
    // 要下线的小程序版本号
    _appVersion   string
    // 模板id
    _templateId   string
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
func (r *TaobaoMiniappTemplateOfflineappRequest) SetClients(_clients []string) error {
    r._clients = _clients
    r.Set("clients", _clients)
    return nil
}

// Clients Getter
func (r TaobaoMiniappTemplateOfflineappRequest) GetClients() []string {
    return r._clients
}
// AppId Setter
// 要下线的小程序app_id
func (r *TaobaoMiniappTemplateOfflineappRequest) SetAppId(_appId string) error {
    r._appId = _appId
    r.Set("app_id", _appId)
    return nil
}

// AppId Getter
func (r TaobaoMiniappTemplateOfflineappRequest) GetAppId() string {
    return r._appId
}
// AppVersion Setter
// 要下线的小程序版本号
func (r *TaobaoMiniappTemplateOfflineappRequest) SetAppVersion(_appVersion string) error {
    r._appVersion = _appVersion
    r.Set("app_version", _appVersion)
    return nil
}

// AppVersion Getter
func (r TaobaoMiniappTemplateOfflineappRequest) GetAppVersion() string {
    return r._appVersion
}
// TemplateId Setter
// 模板id
func (r *TaobaoMiniappTemplateOfflineappRequest) SetTemplateId(_templateId string) error {
    r._templateId = _templateId
    r.Set("template_id", _templateId)
    return nil
}

// TemplateId Getter
func (r TaobaoMiniappTemplateOfflineappRequest) GetTemplateId() string {
    return r._templateId
}