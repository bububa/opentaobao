package auction

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionGovDataMonthlyGetAPIResponse 按月统计法院拍卖数据 API返回值
// taobao.auction.gov.data.monthly.get
//
// 按月统计法院拍卖数据
// 包含：
// 标的件数统计：发布标的件数、结束标的件数、开拍标的件数
// 竞价实况：预计成交金额、出价次数、报名人数
// 在线标的：在线标的件数、意向用户数、网拍围观人次
//
// 最长12个月，月的起始时间不能早于2017年3月
type TaobaoAuctionGovDataMonthlyGetAPIResponse struct {
	model.CommonResponse
	TaobaoAuctionGovDataMonthlyGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAuctionGovDataMonthlyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAuctionGovDataMonthlyGetAPIResponseModel).Reset()
}

// TaobaoAuctionGovDataMonthlyGetAPIResponseModel is 按月统计法院拍卖数据 成功返回结果
type TaobaoAuctionGovDataMonthlyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"auction_gov_data_monthly_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 按月统计结果列表
	Results []CourtsBidStatMonthlyList `json:"results,omitempty" xml:"results>courts_bid_stat_monthly_list,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAuctionGovDataMonthlyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoAuctionGovDataMonthlyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAuctionGovDataMonthlyGetAPIResponse)
	},
}

// GetTaobaoAuctionGovDataMonthlyGetAPIResponse 从 sync.Pool 获取 TaobaoAuctionGovDataMonthlyGetAPIResponse
func GetTaobaoAuctionGovDataMonthlyGetAPIResponse() *TaobaoAuctionGovDataMonthlyGetAPIResponse {
	return poolTaobaoAuctionGovDataMonthlyGetAPIResponse.Get().(*TaobaoAuctionGovDataMonthlyGetAPIResponse)
}

// ReleaseTaobaoAuctionGovDataMonthlyGetAPIResponse 将 TaobaoAuctionGovDataMonthlyGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAuctionGovDataMonthlyGetAPIResponse(v *TaobaoAuctionGovDataMonthlyGetAPIResponse) {
	v.Reset()
	poolTaobaoAuctionGovDataMonthlyGetAPIResponse.Put(v)
}
