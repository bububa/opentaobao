package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallProductUpdateSchemaGetAPIResponse 产品更新规则获取接口 API返回值
// tmall.product.update.schema.get
//
// 获取用户更新产品的规则
type TmallProductUpdateSchemaGetAPIResponse struct {
	model.CommonResponse
	TmallProductUpdateSchemaGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallProductUpdateSchemaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallProductUpdateSchemaGetAPIResponseModel).Reset()
}

// TmallProductUpdateSchemaGetAPIResponseModel is 产品更新规则获取接口 成功返回结果
type TmallProductUpdateSchemaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_product_update_schema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 参数产品ID对产品的更新规则
	UpdateProductSchema string `json:"update_product_schema,omitempty" xml:"update_product_schema,omitempty"`
}

// Reset 清空结构体
func (m *TmallProductUpdateSchemaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UpdateProductSchema = ""
}

var poolTmallProductUpdateSchemaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallProductUpdateSchemaGetAPIResponse)
	},
}

// GetTmallProductUpdateSchemaGetAPIResponse 从 sync.Pool 获取 TmallProductUpdateSchemaGetAPIResponse
func GetTmallProductUpdateSchemaGetAPIResponse() *TmallProductUpdateSchemaGetAPIResponse {
	return poolTmallProductUpdateSchemaGetAPIResponse.Get().(*TmallProductUpdateSchemaGetAPIResponse)
}

// ReleaseTmallProductUpdateSchemaGetAPIResponse 将 TmallProductUpdateSchemaGetAPIResponse 保存到 sync.Pool
func ReleaseTmallProductUpdateSchemaGetAPIResponse(v *TmallProductUpdateSchemaGetAPIResponse) {
	v.Reset()
	poolTmallProductUpdateSchemaGetAPIResponse.Put(v)
}
