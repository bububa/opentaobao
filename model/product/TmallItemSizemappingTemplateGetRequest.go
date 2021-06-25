package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取天猫商品尺码表模板 APIRequest
tmall.item.sizemapping.template.get

获取天猫商品尺码表模板
*/
type TmallItemSizemappingTemplateGetRequest struct {
    model.Params

    // 尺码表模板ID
    templateId   int64 

}

func NewTmallItemSizemappingTemplateGetRequest() *TmallItemSizemappingTemplateGetRequest{
    return &TmallItemSizemappingTemplateGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemSizemappingTemplateGetRequest) GetApiMethodName() string {
    return "tmall.item.sizemapping.template.get"
}

func (r TmallItemSizemappingTemplateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemSizemappingTemplateGetRequest) SetTemplateId(templateId int64) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

func (r TmallItemSizemappingTemplateGetRequest) GetTemplateId() int64 {
    return r.templateId
}

