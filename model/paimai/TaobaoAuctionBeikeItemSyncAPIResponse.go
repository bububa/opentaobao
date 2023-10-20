package paimai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoauctionbeikeitemsyncAPIResponse 贝壳商品同步接口 API返回值
// taobao.auction.beike.item.sync
//
// 贝壳商品同步接口
type TaobaoauctionbeikeitemsyncAPIResponse struct {
	model.CommonResponse
	TaobaoauctionbeikeitemsyncAPIResponseModel
}

// TaobaoauctionbeikeitemsyncAPIResponseModel is 贝壳商品同步接口 成功返回结果
type TaobaoauctionbeikeitemsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"auction_beike_item_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功上传的商品ID列表
	Value []int64 `json:"value,omitempty" xml:"value>int64,omitempty"`
	// 结果描述
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
