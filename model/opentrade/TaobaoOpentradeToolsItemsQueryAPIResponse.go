package opentrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeToolsItemsQueryAPIResponse 交易开放获取商品绑定信息 API返回值
// taobao.opentrade.tools.items.query
//
// 交易开放获取商品绑定信息
type TaobaoOpentradeToolsItemsQueryAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeToolsItemsQueryAPIResponseModel
}

// TaobaoOpentradeToolsItemsQueryAPIResponseModel is 交易开放获取商品绑定信息 成功返回结果
type TaobaoOpentradeToolsItemsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_tools_items_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 已绑定的商品ID列表
	ItemIds []int64 `json:"item_ids,omitempty" xml:"item_ids>int64,omitempty"`
}
