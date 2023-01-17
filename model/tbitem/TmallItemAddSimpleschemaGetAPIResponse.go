package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallItemAddSimpleschemaGetAPIResponse 天猫发布商品规则获取 API返回值
// tmall.item.add.simpleschema.get
//
// 通过商家信息获取商品发布字段和规则。
type TmallItemAddSimpleschemaGetAPIResponse struct {
	model.CommonResponse
	TmallItemAddSimpleschemaGetAPIResponseModel
}

// TmallItemAddSimpleschemaGetAPIResponseModel is 天猫发布商品规则获取 成功返回结果
type TmallItemAddSimpleschemaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_add_simpleschema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回发布商品的规则文档
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
