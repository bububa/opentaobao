package paimai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionBeikeItemSyncAPIResponse 贝壳商品同步接口 API返回值
// taobao.auction.beike.item.sync
//
// 贝壳商品同步接口
type TaobaoAuctionBeikeItemSyncAPIResponse struct {
	model.CommonResponse
	TaobaoAuctionBeikeItemSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAuctionBeikeItemSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAuctionBeikeItemSyncAPIResponseModel).Reset()
}

// TaobaoAuctionBeikeItemSyncAPIResponseModel is 贝壳商品同步接口 成功返回结果
type TaobaoAuctionBeikeItemSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"auction_beike_item_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功上传的商品ID列表
	Value []int64 `json:"value,omitempty" xml:"value>int64,omitempty"`
	// 结果描述
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAuctionBeikeItemSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Value = m.Value[:0]
	m.ResultCode = nil
}

var poolTaobaoAuctionBeikeItemSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAuctionBeikeItemSyncAPIResponse)
	},
}

// GetTaobaoAuctionBeikeItemSyncAPIResponse 从 sync.Pool 获取 TaobaoAuctionBeikeItemSyncAPIResponse
func GetTaobaoAuctionBeikeItemSyncAPIResponse() *TaobaoAuctionBeikeItemSyncAPIResponse {
	return poolTaobaoAuctionBeikeItemSyncAPIResponse.Get().(*TaobaoAuctionBeikeItemSyncAPIResponse)
}

// ReleaseTaobaoAuctionBeikeItemSyncAPIResponse 将 TaobaoAuctionBeikeItemSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAuctionBeikeItemSyncAPIResponse(v *TaobaoAuctionBeikeItemSyncAPIResponse) {
	v.Reset()
	poolTaobaoAuctionBeikeItemSyncAPIResponse.Put(v)
}
