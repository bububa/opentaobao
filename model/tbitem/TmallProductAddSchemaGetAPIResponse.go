package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallProductAddSchemaGetAPIResponse 产品发布规则获取接口 API返回值
// tmall.product.add.schema.get
//
// 获取用户发布产品的规则
type TmallProductAddSchemaGetAPIResponse struct {
	model.CommonResponse
	TmallProductAddSchemaGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallProductAddSchemaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallProductAddSchemaGetAPIResponseModel).Reset()
}

// TmallProductAddSchemaGetAPIResponseModel is 产品发布规则获取接口 成功返回结果
type TmallProductAddSchemaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_product_add_schema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回发布产品的规则文档
	AddProductRule string `json:"add_product_rule,omitempty" xml:"add_product_rule,omitempty"`
}

// Reset 清空结构体
func (m *TmallProductAddSchemaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AddProductRule = ""
}

var poolTmallProductAddSchemaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallProductAddSchemaGetAPIResponse)
	},
}

// GetTmallProductAddSchemaGetAPIResponse 从 sync.Pool 获取 TmallProductAddSchemaGetAPIResponse
func GetTmallProductAddSchemaGetAPIResponse() *TmallProductAddSchemaGetAPIResponse {
	return poolTmallProductAddSchemaGetAPIResponse.Get().(*TmallProductAddSchemaGetAPIResponse)
}

// ReleaseTmallProductAddSchemaGetAPIResponse 将 TmallProductAddSchemaGetAPIResponse 保存到 sync.Pool
func ReleaseTmallProductAddSchemaGetAPIResponse(v *TmallProductAddSchemaGetAPIResponse) {
	v.Reset()
	poolTmallProductAddSchemaGetAPIResponse.Put(v)
}
