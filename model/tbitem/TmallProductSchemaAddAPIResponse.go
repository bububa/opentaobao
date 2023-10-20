package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallProductSchemaAddAPIResponse 使用Schema文件发布一个产品 API返回值
// tmall.product.schema.add
//
// Schema体系发布一个产品
type TmallProductSchemaAddAPIResponse struct {
	model.CommonResponse
	TmallProductSchemaAddAPIResponseModel
}

// Reset 清空结构体
func (m *TmallProductSchemaAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallProductSchemaAddAPIResponseModel).Reset()
}

// TmallProductSchemaAddAPIResponseModel is 使用Schema文件发布一个产品 成功返回结果
type TmallProductSchemaAddAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_product_schema_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 新发产品结果
	AddProductResult string `json:"add_product_result,omitempty" xml:"add_product_result,omitempty"`
}

// Reset 清空结构体
func (m *TmallProductSchemaAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AddProductResult = ""
}

var poolTmallProductSchemaAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallProductSchemaAddAPIResponse)
	},
}

// GetTmallProductSchemaAddAPIResponse 从 sync.Pool 获取 TmallProductSchemaAddAPIResponse
func GetTmallProductSchemaAddAPIResponse() *TmallProductSchemaAddAPIResponse {
	return poolTmallProductSchemaAddAPIResponse.Get().(*TmallProductSchemaAddAPIResponse)
}

// ReleaseTmallProductSchemaAddAPIResponse 将 TmallProductSchemaAddAPIResponse 保存到 sync.Pool
func ReleaseTmallProductSchemaAddAPIResponse(v *TmallProductSchemaAddAPIResponse) {
	v.Reset()
	poolTmallProductSchemaAddAPIResponse.Put(v)
}
