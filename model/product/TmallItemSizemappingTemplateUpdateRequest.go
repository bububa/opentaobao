package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新天猫商品尺码表模板 API请求
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

// 初始化TmallItemSizemappingTemplateUpdateRequest对象
func NewTmallItemSizemappingTemplateUpdateRequest() *TmallItemSizemappingTemplateUpdateRequest{
    return &TmallItemSizemappingTemplateUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemSizemappingTemplateUpdateRequest) GetApiMethodName() string {
    return "tmall.item.sizemapping.template.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemSizemappingTemplateUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TemplateId Setter
// 尺码表模板ID
func (r *TmallItemSizemappingTemplateUpdateRequest) SetTemplateId(templateId int64) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

// TemplateId Getter
func (r TmallItemSizemappingTemplateUpdateRequest) GetTemplateId() int64 {
    return r.templateId
}
// TemplateName Setter
// 尺码表模板名称
func (r *TmallItemSizemappingTemplateUpdateRequest) SetTemplateName(templateName string) error {
    r.templateName = templateName
    r.Set("template_name", templateName)
    return nil
}

// TemplateName Getter
func (r TmallItemSizemappingTemplateUpdateRequest) GetTemplateName() string {
    return r.templateName
}
// TemplateContent Setter
// 尺码表模板内容，格式为"尺码值:维度名称:数值,尺码值:维度名称:数值"。其中，数值的单位，长度单位为厘米（cm），体重单位为公斤（kg）。尺码值，维度数据不能包含数字，特殊字符。数值为0-999.9的数字，且最多一位小数。
func (r *TmallItemSizemappingTemplateUpdateRequest) SetTemplateContent(templateContent string) error {
    r.templateContent = templateContent
    r.Set("template_content", templateContent)
    return nil
}

// TemplateContent Getter
func (r TmallItemSizemappingTemplateUpdateRequest) GetTemplateContent() string {
    return r.templateContent
}
