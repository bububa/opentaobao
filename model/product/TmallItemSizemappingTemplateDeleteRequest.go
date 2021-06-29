package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除天猫商品尺码表模板 API请求
tmall.item.sizemapping.template.delete

删除天猫商品尺码表模板
*/
type TmallItemSizemappingTemplateDeleteRequest struct {
    model.Params
    // 尺码表模板ID
    templateId   int64
}

// 初始化TmallItemSizemappingTemplateDeleteRequest对象
func NewTmallItemSizemappingTemplateDeleteRequest() *TmallItemSizemappingTemplateDeleteRequest{
    return &TmallItemSizemappingTemplateDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemSizemappingTemplateDeleteRequest) GetApiMethodName() string {
    return "tmall.item.sizemapping.template.delete"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemSizemappingTemplateDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TemplateId Setter
// 尺码表模板ID
func (r *TmallItemSizemappingTemplateDeleteRequest) SetTemplateId(templateId int64) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

// TemplateId Getter
func (r TmallItemSizemappingTemplateDeleteRequest) GetTemplateId() int64 {
    return r.templateId
}
