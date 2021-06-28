package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除天猫商品尺码表模板 APIResponse
tmall.item.sizemapping.template.delete

删除天猫商品尺码表模板
*/
type TmallItemSizemappingTemplateDeleteAPIResponse struct {
    model.CommonResponse
    TmallItemSizemappingTemplateDeleteResponse
}

type TmallItemSizemappingTemplateDeleteResponse struct {
    XMLName xml.Name `xml:"tmall_item_sizemapping_template_delete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 尺码表模板ID
    
    TemplateId   int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`

    
}
