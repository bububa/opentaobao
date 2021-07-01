package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallProductSchemaUpdateAPIResponse
产品更新接口 API返回值
tmall.product.schema.update

产品更新接口 */
type TmallProductSchemaUpdateAPIResponse struct {
	model.CommonResponse
	TmallProductSchemaUpdateAPIResponseModel
}

// TmallProductSchemaUpdateAPIResponseModel is 产品更新接口 成功返回结果
type TmallProductSchemaUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_product_schema_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品数据，格式和入参xml_data一致，仅包含产品ID和更新时间
	UpdateProductResult string `json:"update_product_result,omitempty" xml:"update_product_result,omitempty"`
}
