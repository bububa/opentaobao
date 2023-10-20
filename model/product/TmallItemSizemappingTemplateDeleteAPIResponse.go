package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSizemappingTemplateDeleteAPIResponse 删除天猫商品尺码表模板 API返回值
// tmall.item.sizemapping.template.delete
//
// 删除天猫商品尺码表模板
type TmallItemSizemappingTemplateDeleteAPIResponse struct {
	model.CommonResponse
	TmallItemSizemappingTemplateDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemSizemappingTemplateDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemSizemappingTemplateDeleteAPIResponseModel).Reset()
}

// TmallItemSizemappingTemplateDeleteAPIResponseModel is 删除天猫商品尺码表模板 成功返回结果
type TmallItemSizemappingTemplateDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_sizemapping_template_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 尺码表模板ID
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemSizemappingTemplateDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TemplateId = 0
}

var poolTmallItemSizemappingTemplateDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemSizemappingTemplateDeleteAPIResponse)
	},
}

// GetTmallItemSizemappingTemplateDeleteAPIResponse 从 sync.Pool 获取 TmallItemSizemappingTemplateDeleteAPIResponse
func GetTmallItemSizemappingTemplateDeleteAPIResponse() *TmallItemSizemappingTemplateDeleteAPIResponse {
	return poolTmallItemSizemappingTemplateDeleteAPIResponse.Get().(*TmallItemSizemappingTemplateDeleteAPIResponse)
}

// ReleaseTmallItemSizemappingTemplateDeleteAPIResponse 将 TmallItemSizemappingTemplateDeleteAPIResponse 保存到 sync.Pool
func ReleaseTmallItemSizemappingTemplateDeleteAPIResponse(v *TmallItemSizemappingTemplateDeleteAPIResponse) {
	v.Reset()
	poolTmallItemSizemappingTemplateDeleteAPIResponse.Put(v)
}
