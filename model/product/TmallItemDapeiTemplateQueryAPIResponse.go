package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemDapeiTemplateQueryAPIResponse 搭配查询接口 API返回值
// tmall.item.dapei.template.query
//
// 根据条件获取搭配内容
type TmallItemDapeiTemplateQueryAPIResponse struct {
	model.CommonResponse
	TmallItemDapeiTemplateQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemDapeiTemplateQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemDapeiTemplateQueryAPIResponseModel).Reset()
}

// TmallItemDapeiTemplateQueryAPIResponseModel is 搭配查询接口 成功返回结果
type TmallItemDapeiTemplateQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_dapei_template_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallItemDapeiTemplateQueryResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemDapeiTemplateQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallItemDapeiTemplateQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemDapeiTemplateQueryAPIResponse)
	},
}

// GetTmallItemDapeiTemplateQueryAPIResponse 从 sync.Pool 获取 TmallItemDapeiTemplateQueryAPIResponse
func GetTmallItemDapeiTemplateQueryAPIResponse() *TmallItemDapeiTemplateQueryAPIResponse {
	return poolTmallItemDapeiTemplateQueryAPIResponse.Get().(*TmallItemDapeiTemplateQueryAPIResponse)
}

// ReleaseTmallItemDapeiTemplateQueryAPIResponse 将 TmallItemDapeiTemplateQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallItemDapeiTemplateQueryAPIResponse(v *TmallItemDapeiTemplateQueryAPIResponse) {
	v.Reset()
	poolTmallItemDapeiTemplateQueryAPIResponse.Put(v)
}
