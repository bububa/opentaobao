package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallProductSchemaGetAPIResponse
产品信息获取schema获取 API返回值
tmall.product.schema.get

产品信息获取接口schema形式返回 */
type TmallProductSchemaGetAPIResponse struct {
	model.CommonResponse
	TmallProductSchemaGetAPIResponseModel
}

// TmallProductSchemaGetAPIResponseModel is 产品信息获取schema获取 成功返回结果
type TmallProductSchemaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_product_schema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品信息数据。schema形式
	GetProductResult string `json:"get_product_result,omitempty" xml:"get_product_result,omitempty"`
}
