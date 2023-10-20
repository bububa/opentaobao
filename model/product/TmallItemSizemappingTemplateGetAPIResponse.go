package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSizemappingTemplateGetAPIResponse 获取天猫商品尺码表模板 API返回值
// tmall.item.sizemapping.template.get
//
// 获取天猫商品尺码表模板
type TmallItemSizemappingTemplateGetAPIResponse struct {
	model.CommonResponse
	TmallItemSizemappingTemplateGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemSizemappingTemplateGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemSizemappingTemplateGetAPIResponseModel).Reset()
}

// TmallItemSizemappingTemplateGetAPIResponseModel is 获取天猫商品尺码表模板 成功返回结果
type TmallItemSizemappingTemplateGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_sizemapping_template_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 尺码表模板
	SizeMappingTemplate *TmallItemSizemappingTemplateGetModel `json:"size_mapping_template,omitempty" xml:"size_mapping_template,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemSizemappingTemplateGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SizeMappingTemplate = nil
}

var poolTmallItemSizemappingTemplateGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemSizemappingTemplateGetAPIResponse)
	},
}

// GetTmallItemSizemappingTemplateGetAPIResponse 从 sync.Pool 获取 TmallItemSizemappingTemplateGetAPIResponse
func GetTmallItemSizemappingTemplateGetAPIResponse() *TmallItemSizemappingTemplateGetAPIResponse {
	return poolTmallItemSizemappingTemplateGetAPIResponse.Get().(*TmallItemSizemappingTemplateGetAPIResponse)
}

// ReleaseTmallItemSizemappingTemplateGetAPIResponse 将 TmallItemSizemappingTemplateGetAPIResponse 保存到 sync.Pool
func ReleaseTmallItemSizemappingTemplateGetAPIResponse(v *TmallItemSizemappingTemplateGetAPIResponse) {
	v.Reset()
	poolTmallItemSizemappingTemplateGetAPIResponse.Put(v)
}
