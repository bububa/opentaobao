package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取天猫商品尺码表模板 APIResponse
tmall.item.sizemapping.template.get

获取天猫商品尺码表模板
*/
type TmallItemSizemappingTemplateGetAPIResponse struct {
    model.CommonResponse
    Response *TmallItemSizemappingTemplateGetResponse `json:"tmall_item_sizemapping_template_get_response,omitempty"`
}

type TmallItemSizemappingTemplateGetResponse struct {

    // 尺码表模板
    SizeMappingTemplate   *TmallItemSizemappingTemplateGetModel `json:"size_mapping_template,omitempty"`

}
