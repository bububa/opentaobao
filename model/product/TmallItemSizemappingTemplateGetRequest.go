package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取天猫商品尺码表模板 API请求
tmall.item.sizemapping.template.get

获取天猫商品尺码表模板
*/
type TmallItemSizemappingTemplateGetRequest struct {
    model.Params
    // 尺码表模板ID
    templateId   int64
}

// 初始化TmallItemSizemappingTemplateGetRequest对象
func NewTmallItemSizemappingTemplateGetRequest() *TmallItemSizemappingTemplateGetRequest{
    return &TmallItemSizemappingTemplateGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemSizemappingTemplateGetRequest) GetApiMethodName() string {
    return "tmall.item.sizemapping.template.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemSizemappingTemplateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TemplateId Setter
// 尺码表模板ID
func (r *TmallItemSizemappingTemplateGetRequest) SetTemplateId(templateId int64) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

// TemplateId Getter
func (r TmallItemSizemappingTemplateGetRequest) GetTemplateId() int64 {
    return r.templateId
}
