package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallItemUpdateSchemaGetAPIResponse 天猫编辑商品规则获取 API返回值
// tmall.item.update.schema.get
//
// Schema方式编辑天猫商品时，编辑商品规则获取
type TmallItemUpdateSchemaGetAPIResponse struct {
	model.CommonResponse
	TmallItemUpdateSchemaGetAPIResponseModel
}

// TmallItemUpdateSchemaGetAPIResponseModel is 天猫编辑商品规则获取 成功返回结果
type TmallItemUpdateSchemaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_update_schema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回发布商品的规则文档
	UpdateItemResult string `json:"update_item_result,omitempty" xml:"update_item_result,omitempty"`
}
