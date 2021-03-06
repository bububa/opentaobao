package openmall

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallItemsQueryAPIResponse 批量获取商品列表 API返回值
// taobao.openmall.items.query
//
// 批量获取对联盟开放的商品列表。
type TaobaoOpenmallItemsQueryAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallItemsQueryAPIResponseModel
}

// TaobaoOpenmallItemsQueryAPIResponseModel is 批量获取商品列表 成功返回结果
type TaobaoOpenmallItemsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_items_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoOpenmallItemsQueryResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
