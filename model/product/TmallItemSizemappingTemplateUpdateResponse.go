package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新天猫商品尺码表模板 APIResponse
tmall.item.sizemapping.template.update

更新天猫商品尺码表模板
*/
type TmallItemSizemappingTemplateUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemSizemappingTemplateUpdateResponse `json:"tmall_item_sizemapping_template_update_response,omitempty"` 
    TmallItemSizemappingTemplateUpdateResponse
}

/* model for simplify = false
type TmallItemSizemappingTemplateUpdateResponse struct {

    // 尺码表模板
    
    SizeMappingTemplate  *struct {
        SizeMappingTemplateDo  *SizeMappingTemplateDo `json:"size_mapping_template_do,omitempty"`
    } `json:"size_mapping_template,omitempty"`
    

}
*/

type TmallItemSizemappingTemplateUpdateResponse struct {

    // 尺码表模板
    SizeMappingTemplate   *SizeMappingTemplateDo `json:"size_mapping_template,omitempty"`

}
