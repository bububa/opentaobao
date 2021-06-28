package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取天猫商品尺码表模板列表 APIResponse
tmall.item.sizemapping.templates.list

获取所有尺码表模板列表。
*/
type TmallItemSizemappingTemplatesListAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemSizemappingTemplatesListResponse `json:"tmall_item_sizemapping_templates_list_response,omitempty"` 
    TmallItemSizemappingTemplatesListResponse
}

/* model for simplify = false
type TmallItemSizemappingTemplatesListResponse struct {

    // 尺码表模板列表
    
    SizeMappingTemplates  struct {
        SizeMappingTemplate  []SizeMappingTemplate `json:"size_mapping_template,omitempty"`
    } `json:"size_mapping_templates,omitempty"`
    

}
*/

type TmallItemSizemappingTemplatesListResponse struct {

    // 尺码表模板列表
    SizeMappingTemplates   []SizeMappingTemplate `json:"size_mapping_templates,omitempty"`

}
