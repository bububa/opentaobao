package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSchemaAddAPIResponse 天猫根据规则发布商品 API返回值
// tmall.item.schema.add
//
// 天猫TopSchema发布商品。
type TmallItemSchemaAddAPIResponse struct {
	model.CommonResponse
	TmallItemSchemaAddAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemSchemaAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemSchemaAddAPIResponseModel).Reset()
}

// TmallItemSchemaAddAPIResponseModel is 天猫根据规则发布商品 成功返回结果
type TmallItemSchemaAddAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_schema_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回商品发布结果
	AddItemResult string `json:"add_item_result,omitempty" xml:"add_item_result,omitempty"`
	// 发布商品操作成功时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemSchemaAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AddItemResult = ""
	m.GmtCreate = ""
}

var poolTmallItemSchemaAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemSchemaAddAPIResponse)
	},
}

// GetTmallItemSchemaAddAPIResponse 从 sync.Pool 获取 TmallItemSchemaAddAPIResponse
func GetTmallItemSchemaAddAPIResponse() *TmallItemSchemaAddAPIResponse {
	return poolTmallItemSchemaAddAPIResponse.Get().(*TmallItemSchemaAddAPIResponse)
}

// ReleaseTmallItemSchemaAddAPIResponse 将 TmallItemSchemaAddAPIResponse 保存到 sync.Pool
func ReleaseTmallItemSchemaAddAPIResponse(v *TmallItemSchemaAddAPIResponse) {
	v.Reset()
	poolTmallItemSchemaAddAPIResponse.Put(v)
}
