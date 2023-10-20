package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSizemappingTemplateUpdateAPIResponse 更新天猫商品尺码表模板 API返回值
// tmall.item.sizemapping.template.update
//
// 更新天猫商品尺码表模板
type TmallItemSizemappingTemplateUpdateAPIResponse struct {
	model.CommonResponse
	TmallItemSizemappingTemplateUpdateAPIResponseModel
}

// TmallItemSizemappingTemplateUpdateAPIResponseModel is 更新天猫商品尺码表模板 成功返回结果
type TmallItemSizemappingTemplateUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_sizemapping_template_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 尺码表模板
	SizeMappingTemplate *SizeMappingTemplateDo `json:"size_mapping_template,omitempty" xml:"size_mapping_template,omitempty"`
}
