package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除运费模板 APIRequest
taobao.delivery.template.delete

根据用户指定的模板ID删除指定的模板
*/
type TaobaoDeliveryTemplateDeleteRequest struct {
    model.Params

    // 运费模板ID
    templateId   int64 

}

func NewTaobaoDeliveryTemplateDeleteRequest() *TaobaoDeliveryTemplateDeleteRequest{
    return &TaobaoDeliveryTemplateDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDeliveryTemplateDeleteRequest) GetApiMethodName() string {
    return "taobao.delivery.template.delete"
}

func (r TaobaoDeliveryTemplateDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoDeliveryTemplateDeleteRequest) SetTemplateId(templateId int64) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

func (r TaobaoDeliveryTemplateDeleteRequest) GetTemplateId() int64 {
    return r.templateId
}

