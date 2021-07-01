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
type TmallItemSizemappingTemplateGetAPIRequest struct {
    model.Params
    // 尺码表模板ID
    _templateId   int64
}

// 初始化TmallItemSizemappingTemplateGetAPIRequest对象
func NewTmallItemSizemappingTemplateGetRequest() *TmallItemSizemappingTemplateGetAPIRequest{
    return &TmallItemSizemappingTemplateGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemSizemappingTemplateGetAPIRequest) GetApiMethodName() string {
    return "tmall.item.sizemapping.template.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemSizemappingTemplateGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TemplateId Setter
// 尺码表模板ID
func (r *TmallItemSizemappingTemplateGetAPIRequest) SetTemplateId(_templateId int64) error {
    r._templateId = _templateId
    r.Set("template_id", _templateId)
    return nil
}

// TemplateId Getter
func (r TmallItemSizemappingTemplateGetAPIRequest) GetTemplateId() int64 {
    return r._templateId
}
