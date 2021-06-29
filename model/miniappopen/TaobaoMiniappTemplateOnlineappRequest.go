package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上线实例化应用 APIRequest
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

func NewTaobaoMiniappTemplateOnlineappRequest() *TaobaoMiniappTemplateOnlineappRequest{
    return &TaobaoMiniappTemplateOnlineappRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappTemplateOnlineappRequest) GetApiMethodName() string {
    return "taobao.miniapp.template.onlineapp"
}

func (r TaobaoMiniappTemplateOnlineappRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappTemplateOnlineappRequest) SetClients(clients []string) error {
    r.clients = clients
    r.Set("clients", clients)
    return nil
}

func (r TaobaoMiniappTemplateOnlineappRequest) GetClients() []string {
    return r.clients
}

func (r *TaobaoMiniappTemplateOnlineappRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

func (r TaobaoMiniappTemplateOnlineappRequest) GetAppId() string {
    return r.appId
}

func (r *TaobaoMiniappTemplateOnlineappRequest) SetTemplateId(templateId string) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

func (r TaobaoMiniappTemplateOnlineappRequest) GetTemplateId() string {
    return r.templateId
}

func (r *TaobaoMiniappTemplateOnlineappRequest) SetTemplateVersion(templateVersion string) error {
    r.templateVersion = templateVersion
    r.Set("template_version", templateVersion)
    return nil
}

func (r TaobaoMiniappTemplateOnlineappRequest) GetTemplateVersion() string {
    return r.templateVersion
}

func (r *TaobaoMiniappTemplateOnlineappRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

func (r TaobaoMiniappTemplateOnlineappRequest) GetAppVersion() string {
    return r.appVersion
}

