package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallProductSchemaUpdateAPIResponse 产品更新接口 API返回值
// tmall.product.schema.update
//
// 产品更新接口
type TmallProductSchemaUpdateAPIResponse struct {
	model.CommonResponse
	TmallProductSchemaUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallProductSchemaUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallProductSchemaUpdateAPIResponseModel).Reset()
}

// TmallProductSchemaUpdateAPIResponseModel is 产品更新接口 成功返回结果
type TmallProductSchemaUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_product_schema_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品数据，格式和入参xml_data一致，仅包含产品ID和更新时间
	UpdateProductResult string `json:"update_product_result,omitempty" xml:"update_product_result,omitempty"`
}

// Reset 清空结构体
func (m *TmallProductSchemaUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UpdateProductResult = ""
}

var poolTmallProductSchemaUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallProductSchemaUpdateAPIResponse)
	},
}

// GetTmallProductSchemaUpdateAPIResponse 从 sync.Pool 获取 TmallProductSchemaUpdateAPIResponse
func GetTmallProductSchemaUpdateAPIResponse() *TmallProductSchemaUpdateAPIResponse {
	return poolTmallProductSchemaUpdateAPIResponse.Get().(*TmallProductSchemaUpdateAPIResponse)
}

// ReleaseTmallProductSchemaUpdateAPIResponse 将 TmallProductSchemaUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallProductSchemaUpdateAPIResponse(v *TmallProductSchemaUpdateAPIResponse) {
	v.Reset()
	poolTmallProductSchemaUpdateAPIResponse.Put(v)
}
