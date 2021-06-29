package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除运费模板 API请求
taobao.delivery.template.delete

根据用户指定的模板ID删除指定的模板
*/
type TaobaoDeliveryTemplateDeleteRequest struct {
    model.Params
    // 运费模板ID
    templateId   int64
}

// 初始化TaobaoDeliveryTemplateDeleteRequest对象
func NewTaobaoDeliveryTemplateDeleteRequest() *TaobaoDeliveryTemplateDeleteRequest{
    return &TaobaoDeliveryTemplateDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDeliveryTemplateDeleteRequest) GetApiMethodName() string {
    return "taobao.delivery.template.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDeliveryTemplateDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TemplateId Setter
// 运费模板ID
func (r *TaobaoDeliveryTemplateDeleteRequest) SetTemplateId(templateId int64) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

// TemplateId Getter
func (r TaobaoDeliveryTemplateDeleteRequest) GetTemplateId() int64 {
    return r.templateId
}
