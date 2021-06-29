package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新天猫商品尺码表模板 APIResponse
tmall.item.sizemapping.template.update

更新天猫商品尺码表模板
*/
type TmallItemSizemappingTemplateUpdateAPIResponse struct {
    model.CommonResponse
    TmallItemSizemappingTemplateUpdateResponse
}

type TmallItemSizemappingTemplateUpdateResponse struct {
    XMLName xml.Name `xml:"tmall_item_sizemapping_template_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 尺码表模板
    
    SizeMappingTemplate   *SizeMappingTemplateDo `json:"size_mapping_template,omitempty" xml:"size_mapping_template,omitempty"`

    
}
