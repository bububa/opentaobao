package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新实例化应用 APIRequest
taobao.miniapp.template.updateapp

商家应用c端模板实例化小程序更新，生成新的版本，但不会自动上线新版本
*/
type TaobaoMiniappTemplateUpdateappRequest struct {
    model.Params

    // 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
    clients   []string 

    // 应用id，如果应用在对应端上已有的最新版本所使用的模板id,模板version和extjson，与本次入参一致，则认为不需要更新，返回已有的版本。
    appId   string 

    // 扩展信息。线上版本使用的template_id与传入的template_id一致时，可不填。
    extJson   string 

    // 模板id
    templateId   string 

    // 模板版本
    templateVersion   string 

    // Logo更新，1月5次
    icon   string 

    // 描述更新，1年5次
    desc   string 

    // 简称更新，1年5次
    alias   string 

}

func NewTaobaoMiniappTemplateUpdateappRequest() *TaobaoMiniappTemplateUpdateappRequest{
    return &TaobaoMiniappTemplateUpdateappRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappTemplateUpdateappRequest) GetApiMethodName() string {
    return "taobao.miniapp.template.updateapp"
}

func (r TaobaoMiniappTemplateUpdateappRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappTemplateUpdateappRequest) SetClients(clients []string) error {
    r.clients = clients
    r.Set("clients", clients)
    return nil
}

func (r TaobaoMiniappTemplateUpdateappRequest) GetClients() []string {
    return r.clients
}

func (r *TaobaoMiniappTemplateUpdateappRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

func (r TaobaoMiniappTemplateUpdateappRequest) GetAppId() string {
    return r.appId
}

func (r *TaobaoMiniappTemplateUpdateappRequest) SetExtJson(extJson string) error {
    r.extJson = extJson
    r.Set("ext_json", extJson)
    return nil
}

func (r TaobaoMiniappTemplateUpdateappRequest) GetExtJson() string {
    return r.extJson
}

func (r *TaobaoMiniappTemplateUpdateappRequest) SetTemplateId(templateId string) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

func (r TaobaoMiniappTemplateUpdateappRequest) GetTemplateId() string {
    return r.templateId
}

func (r *TaobaoMiniappTemplateUpdateappRequest) SetTemplateVersion(templateVersion string) error {
    r.templateVersion = templateVersion
    r.Set("template_version", templateVersion)
    return nil
}

func (r TaobaoMiniappTemplateUpdateappRequest) GetTemplateVersion() string {
    return r.templateVersion
}

func (r *TaobaoMiniappTemplateUpdateappRequest) SetIcon(icon string) error {
    r.icon = icon
    r.Set("icon", icon)
    return nil
}

func (r TaobaoMiniappTemplateUpdateappRequest) GetIcon() string {
    return r.icon
}

func (r *TaobaoMiniappTemplateUpdateappRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

func (r TaobaoMiniappTemplateUpdateappRequest) GetDesc() string {
    return r.desc
}

func (r *TaobaoMiniappTemplateUpdateappRequest) SetAlias(alias string) error {
    r.alias = alias
    r.Set("alias", alias)
    return nil
}

func (r TaobaoMiniappTemplateUpdateappRequest) GetAlias() string {
    return r.alias
}

