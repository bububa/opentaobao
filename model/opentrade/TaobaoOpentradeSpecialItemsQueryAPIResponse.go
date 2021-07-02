package opentrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeSpecialItemsQueryAPIResponse 专属下单获取商品绑定信息 API返回值
// taobao.opentrade.special.items.query
//
// 专属下单获取商品绑定信息
type TaobaoOpentradeSpecialItemsQueryAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeSpecialItemsQueryAPIResponseModel
}

// TaobaoOpentradeSpecialItemsQueryAPIResponseModel is 专属下单获取商品绑定信息 成功返回结果
type TaobaoOpentradeSpecialItemsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_special_items_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 已绑定的商品ID列表
	Items []int64 `json:"items,omitempty" xml:"items>int64,omitempty"`
}
