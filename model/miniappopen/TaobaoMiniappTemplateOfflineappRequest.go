package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
下线实例化应用 APIRequest
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

func NewTaobaoMiniappTemplateOfflineappRequest() *TaobaoMiniappTemplateOfflineappRequest{
    return &TaobaoMiniappTemplateOfflineappRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappTemplateOfflineappRequest) GetApiMethodName() string {
    return "taobao.miniapp.template.offlineapp"
}

func (r TaobaoMiniappTemplateOfflineappRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappTemplateOfflineappRequest) SetClients(clients []string) error {
    r.clients = clients
    r.Set("clients", clients)
    return nil
}

func (r TaobaoMiniappTemplateOfflineappRequest) GetClients() []string {
    return r.clients
}

func (r *TaobaoMiniappTemplateOfflineappRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

func (r TaobaoMiniappTemplateOfflineappRequest) GetAppId() string {
    return r.appId
}

func (r *TaobaoMiniappTemplateOfflineappRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

func (r TaobaoMiniappTemplateOfflineappRequest) GetAppVersion() string {
    return r.appVersion
}

func (r *TaobaoMiniappTemplateOfflineappRequest) SetTemplateId(templateId string) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

func (r TaobaoMiniappTemplateOfflineappRequest) GetTemplateId() string {
    return r.templateId
}

