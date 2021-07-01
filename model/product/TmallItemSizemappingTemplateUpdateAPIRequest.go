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
type TmallItemSizemappingTemplateUpdateAPIRequest struct {
    model.Params
    // 尺码表模板ID
    _templateId   int64
    // 尺码表模板名称
    _templateName   string
    // 尺码表模板内容，格式为"尺码值:维度名称:数值,尺码值:维度名称:数值"。其中，数值的单位，长度单位为厘米（cm），体重单位为公斤（kg）。尺码值，维度数据不能包含数字，特殊字符。数值为0-999.9的数字，且最多一位小数。
    _templateContent   string
}

// 初始化TmallItemSizemappingTemplateUpdateAPIRequest对象
func NewTmallItemSizemappingTemplateUpdateRequest() *TmallItemSizemappingTemplateUpdateAPIRequest{
    return &TmallItemSizemappingTemplateUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemSizemappingTemplateUpdateAPIRequest) GetApiMethodName() string {
    return "tmall.item.sizemapping.template.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemSizemappingTemplateUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TemplateId Setter
// 尺码表模板ID
func (r *TmallItemSizemappingTemplateUpdateAPIRequest) SetTemplateId(_templateId int64) error {
    r._templateId = _templateId
    r.Set("template_id", _templateId)
    return nil
}

// TemplateId Getter
func (r TmallItemSizemappingTemplateUpdateAPIRequest) GetTemplateId() int64 {
    return r._templateId
}
// TemplateName Setter
// 尺码表模板名称
func (r *TmallItemSizemappingTemplateUpdateAPIRequest) SetTemplateName(_templateName string) error {
    r._templateName = _templateName
    r.Set("template_name", _templateName)
    return nil
}

// TemplateName Getter
func (r TmallItemSizemappingTemplateUpdateAPIRequest) GetTemplateName() string {
    return r._templateName
}
// TemplateContent Setter
// 尺码表模板内容，格式为"尺码值:维度名称:数值,尺码值:维度名称:数值"。其中，数值的单位，长度单位为厘米（cm），体重单位为公斤（kg）。尺码值，维度数据不能包含数字，特殊字符。数值为0-999.9的数字，且最多一位小数。
func (r *TmallItemSizemappingTemplateUpdateAPIRequest) SetTemplateContent(_templateContent string) error {
    r._templateContent = _templateContent
    r.Set("template_content", _templateContent)
    return nil
}

// TemplateContent Getter
func (r TmallItemSizemappingTemplateUpdateAPIRequest) GetTemplateContent() string {
    return r._templateContent
}
