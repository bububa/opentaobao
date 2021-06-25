package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新天猫商品尺码表模板 APIRequest
tmall.item.sizemapping.template.update

更新天猫商品尺码表模板
*/
type TmallItemSizemappingTemplateUpdateRequest struct {
    model.Params

    // 尺码表模板ID
    templateId   int64 

    // 尺码表模板名称
    templateName   string 

    // 尺码表模板内容，格式为"尺码值:维度名称:数值,尺码值:维度名称:数值"。其中，数值的单位，长度单位为厘米（cm），体重单位为公斤（kg）。尺码值，维度数据不能包含数字，特殊字符。数值为0-999.9的数字，且最多一位小数。
    templateContent   string 

}

func NewTmallItemSizemappingTemplateUpdateRequest() *TmallItemSizemappingTemplateUpdateRequest{
    return &TmallItemSizemappingTemplateUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemSizemappingTemplateUpdateRequest) GetApiMethodName() string {
    return "tmall.item.sizemapping.template.update"
}

func (r TmallItemSizemappingTemplateUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemSizemappingTemplateUpdateRequest) SetTemplateId(templateId int64) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

func (r TmallItemSizemappingTemplateUpdateRequest) GetTemplateId() int64 {
    return r.templateId
}

func (r *TmallItemSizemappingTemplateUpdateRequest) SetTemplateName(templateName string) error {
    r.templateName = templateName
    r.Set("template_name", templateName)
    return nil
}

func (r TmallItemSizemappingTemplateUpdateRequest) GetTemplateName() string {
    return r.templateName
}

func (r *TmallItemSizemappingTemplateUpdateRequest) SetTemplateContent(templateContent string) error {
    r.templateContent = templateContent
    r.Set("template_content", templateContent)
    return nil
}

func (r TmallItemSizemappingTemplateUpdateRequest) GetTemplateContent() string {
    return r.templateContent
}

