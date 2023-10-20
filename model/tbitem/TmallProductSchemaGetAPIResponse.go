package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallProductSchemaGetAPIResponse 产品信息获取schema获取 API返回值
// tmall.product.schema.get
//
// 产品信息获取接口schema形式返回
type TmallProductSchemaGetAPIResponse struct {
	model.CommonResponse
	TmallProductSchemaGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallProductSchemaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallProductSchemaGetAPIResponseModel).Reset()
}

// TmallProductSchemaGetAPIResponseModel is 产品信息获取schema获取 成功返回结果
type TmallProductSchemaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_product_schema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品信息数据。schema形式
	GetProductResult string `json:"get_product_result,omitempty" xml:"get_product_result,omitempty"`
}

// Reset 清空结构体
func (m *TmallProductSchemaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.GetProductResult = ""
}

var poolTmallProductSchemaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallProductSchemaGetAPIResponse)
	},
}

// GetTmallProductSchemaGetAPIResponse 从 sync.Pool 获取 TmallProductSchemaGetAPIResponse
func GetTmallProductSchemaGetAPIResponse() *TmallProductSchemaGetAPIResponse {
	return poolTmallProductSchemaGetAPIResponse.Get().(*TmallProductSchemaGetAPIResponse)
}

// ReleaseTmallProductSchemaGetAPIResponse 将 TmallProductSchemaGetAPIResponse 保存到 sync.Pool
func ReleaseTmallProductSchemaGetAPIResponse(v *TmallProductSchemaGetAPIResponse) {
	v.Reset()
	poolTmallProductSchemaGetAPIResponse.Put(v)
}
