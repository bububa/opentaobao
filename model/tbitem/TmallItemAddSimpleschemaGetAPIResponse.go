package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemAddSimpleschemaGetAPIResponse 天猫发布商品规则获取 API返回值
// tmall.item.add.simpleschema.get
//
// 通过商家信息获取商品发布字段和规则。
type TmallItemAddSimpleschemaGetAPIResponse struct {
	model.CommonResponse
	TmallItemAddSimpleschemaGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemAddSimpleschemaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemAddSimpleschemaGetAPIResponseModel).Reset()
}

// TmallItemAddSimpleschemaGetAPIResponseModel is 天猫发布商品规则获取 成功返回结果
type TmallItemAddSimpleschemaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_add_simpleschema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回发布商品的规则文档
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemAddSimpleschemaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTmallItemAddSimpleschemaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemAddSimpleschemaGetAPIResponse)
	},
}

// GetTmallItemAddSimpleschemaGetAPIResponse 从 sync.Pool 获取 TmallItemAddSimpleschemaGetAPIResponse
func GetTmallItemAddSimpleschemaGetAPIResponse() *TmallItemAddSimpleschemaGetAPIResponse {
	return poolTmallItemAddSimpleschemaGetAPIResponse.Get().(*TmallItemAddSimpleschemaGetAPIResponse)
}

// ReleaseTmallItemAddSimpleschemaGetAPIResponse 将 TmallItemAddSimpleschemaGetAPIResponse 保存到 sync.Pool
func ReleaseTmallItemAddSimpleschemaGetAPIResponse(v *TmallItemAddSimpleschemaGetAPIResponse) {
	v.Reset()
	poolTmallItemAddSimpleschemaGetAPIResponse.Put(v)
}
