package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除天猫商品尺码表模板 APIRequest
tmall.item.sizemapping.template.delete

删除天猫商品尺码表模板
*/
type TmallItemSizemappingTemplateDeleteRequest struct {
    model.Params

    // 尺码表模板ID
    templateId   int64 

}

func NewTmallItemSizemappingTemplateDeleteRequest() *TmallItemSizemappingTemplateDeleteRequest{
    return &TmallItemSizemappingTemplateDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemSizemappingTemplateDeleteRequest) GetApiMethodName() string {
    return "tmall.item.sizemapping.template.delete"
}

func (r TmallItemSizemappingTemplateDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemSizemappingTemplateDeleteRequest) SetTemplateId(templateId int64) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

func (r TmallItemSizemappingTemplateDeleteRequest) GetTemplateId() int64 {
    return r.templateId
}

