package product

import (
	"encoding/xml"

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
