package auction

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionGovDataTopnGetAPIResponse 根据不同维度，获取排行榜列表 API返回值
// taobao.auction.gov.data.topn.get
//
// 根据不同时间维度(周,月,年)，获取(成交额或发拍件数)排行榜列表
type TaobaoAuctionGovDataTopnGetAPIResponse struct {
	model.CommonResponse
	TaobaoAuctionGovDataTopnGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAuctionGovDataTopnGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAuctionGovDataTopnGetAPIResponseModel).Reset()
}

// TaobaoAuctionGovDataTopnGetAPIResponseModel is 根据不同维度，获取排行榜列表 成功返回结果
type TaobaoAuctionGovDataTopnGetAPIResponseModel struct {
	XMLName xml.Name `xml:"auction_gov_data_topn_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 法院维度标的统计排行
	Ranks []CourtsBidStatTopnDto `json:"ranks,omitempty" xml:"ranks>courts_bid_stat_topn_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAuctionGovDataTopnGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Ranks = m.Ranks[:0]
}

var poolTaobaoAuctionGovDataTopnGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAuctionGovDataTopnGetAPIResponse)
	},
}

// GetTaobaoAuctionGovDataTopnGetAPIResponse 从 sync.Pool 获取 TaobaoAuctionGovDataTopnGetAPIResponse
func GetTaobaoAuctionGovDataTopnGetAPIResponse() *TaobaoAuctionGovDataTopnGetAPIResponse {
	return poolTaobaoAuctionGovDataTopnGetAPIResponse.Get().(*TaobaoAuctionGovDataTopnGetAPIResponse)
}

// ReleaseTaobaoAuctionGovDataTopnGetAPIResponse 将 TaobaoAuctionGovDataTopnGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAuctionGovDataTopnGetAPIResponse(v *TaobaoAuctionGovDataTopnGetAPIResponse) {
	v.Reset()
	poolTaobaoAuctionGovDataTopnGetAPIResponse.Put(v)
}
