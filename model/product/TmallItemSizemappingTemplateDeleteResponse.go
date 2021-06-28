package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除天猫商品尺码表模板 APIResponse
tmall.item.sizemapping.template.delete

删除天猫商品尺码表模板
*/
type TmallItemSizemappingTemplateDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemSizemappingTemplateDeleteResponse `json:"tmall_item_sizemapping_template_delete_response,omitempty"` 
    TmallItemSizemappingTemplateDeleteResponse
}

/* model for simplify = false
type TmallItemSizemappingTemplateDeleteResponse struct {

    // 尺码表模板ID
    
    TemplateId   int64 `json:"template_id,omitempty"`
    

}
*/

type TmallItemSizemappingTemplateDeleteResponse struct {

    // 尺码表模板ID
    TemplateId   int64 `json:"template_id,omitempty"`

}
