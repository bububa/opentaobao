package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回滚实例化应用 APIRequest
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

func NewTaobaoMiniappTemplateRollbackRequest() *TaobaoMiniappTemplateRollbackRequest{
    return &TaobaoMiniappTemplateRollbackRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappTemplateRollbackRequest) GetApiMethodName() string {
    return "taobao.miniapp.template.rollback"
}

func (r TaobaoMiniappTemplateRollbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappTemplateRollbackRequest) SetClients(clients []string) error {
    r.clients = clients
    r.Set("clients", clients)
    return nil
}

func (r TaobaoMiniappTemplateRollbackRequest) GetClients() []string {
    return r.clients
}

func (r *TaobaoMiniappTemplateRollbackRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

func (r TaobaoMiniappTemplateRollbackRequest) GetAppId() string {
    return r.appId
}

func (r *TaobaoMiniappTemplateRollbackRequest) SetAppVersion(appVersion string) error {
    r.appVersion = appVersion
    r.Set("app_version", appVersion)
    return nil
}

func (r TaobaoMiniappTemplateRollbackRequest) GetAppVersion() string {
    return r.appVersion
}

func (r *TaobaoMiniappTemplateRollbackRequest) SetTemplateId(templateId string) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

func (r TaobaoMiniappTemplateRollbackRequest) GetTemplateId() string {
    return r.templateId
}

func (r *TaobaoMiniappTemplateRollbackRequest) SetTemplateVersion(templateVersion string) error {
    r.templateVersion = templateVersion
    r.Set("template_version", templateVersion)
    return nil
}

func (r TaobaoMiniappTemplateRollbackRequest) GetTemplateVersion() string {
    return r.templateVersion
}

