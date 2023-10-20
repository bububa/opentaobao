package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallitemsizemappingtemplatedeleteAPIResponse 删除天猫商品尺码表模板 API返回值
// tmall.item.sizemapping.template.delete
//
// 删除天猫商品尺码表模板
type TmallitemsizemappingtemplatedeleteAPIResponse struct {
	model.CommonResponse
	TmallitemsizemappingtemplatedeleteAPIResponseModel
}

// TmallitemsizemappingtemplatedeleteAPIResponseModel is 删除天猫商品尺码表模板 成功返回结果
type TmallitemsizemappingtemplatedeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_sizemapping_template_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 尺码表模板ID
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
}
