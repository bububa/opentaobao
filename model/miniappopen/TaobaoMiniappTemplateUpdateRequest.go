package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（已废弃）更新实例化应用 APIRequest
taobao.miniapp.template.update

商家应用c端模板实例化小程序更新
*/
type TaobaoMiniappTemplateUpdateRequest struct {
    model.Params

    // 要更新的投放端,目前可投放： taobao(淘宝),tmall(天猫)
    clients   []string 

    // 应用id
    id   string 

    // schema信息，不填且 应用线上版本使用的templateId与传入的templateId不一致，则会报错; 一致，则复用线上版本的schema。
    extJson   string 

    // 模板id
    templateId   string 

    // 模板版本
    templateVersion   string 

}

func NewTaobaoMiniappTemplateUpdateRequest() *TaobaoMiniappTemplateUpdateRequest{
    return &TaobaoMiniappTemplateUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappTemplateUpdateRequest) GetApiMethodName() string {
    return "taobao.miniapp.template.update"
}

func (r TaobaoMiniappTemplateUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappTemplateUpdateRequest) SetClients(clients []string) error {
    r.clients = clients
    r.Set("clients", clients)
    return nil
}

func (r TaobaoMiniappTemplateUpdateRequest) GetClients() []string {
    return r.clients
}

func (r *TaobaoMiniappTemplateUpdateRequest) SetId(id string) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r TaobaoMiniappTemplateUpdateRequest) GetId() string {
    return r.id
}

func (r *TaobaoMiniappTemplateUpdateRequest) SetExtJson(extJson string) error {
    r.extJson = extJson
    r.Set("ext_json", extJson)
    return nil
}

func (r TaobaoMiniappTemplateUpdateRequest) GetExtJson() string {
    return r.extJson
}

func (r *TaobaoMiniappTemplateUpdateRequest) SetTemplateId(templateId string) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

func (r TaobaoMiniappTemplateUpdateRequest) GetTemplateId() string {
    return r.templateId
}

func (r *TaobaoMiniappTemplateUpdateRequest) SetTemplateVersion(templateVersion string) error {
    r.templateVersion = templateVersion
    r.Set("template_version", templateVersion)
    return nil
}

func (r TaobaoMiniappTemplateUpdateRequest) GetTemplateVersion() string {
    return r.templateVersion
}

