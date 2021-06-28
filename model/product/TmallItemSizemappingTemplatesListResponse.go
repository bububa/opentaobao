package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取天猫商品尺码表模板列表 APIResponse
tmall.item.sizemapping.templates.list

获取所有尺码表模板列表。
*/
type TmallItemSizemappingTemplatesListAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_item_sizemapping_templates_list_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 尺码表模板列表
    
    SizeMappingTemplates   []SizeMappingTemplate `json:"size_mapping_templates,omitempty" xml:"