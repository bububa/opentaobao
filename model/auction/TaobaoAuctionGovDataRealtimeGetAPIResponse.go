package auction

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionGovDataRealtimeGetAPIResponse 获取实时(今日)统计数据 API返回值
// taobao.auction.gov.data.realtime.get
//
// 提供查询当日法院及下属法院的拍卖统计数据
type TaobaoAuctionGovDataRealtimeGetAPIResponse struct {
	model.CommonResponse
	TaobaoAuctionGovDataRealtimeGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAuctionGovDataRealtimeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAuctionGovDataRealtimeGetAPIResponseModel).Reset()
}

// TaobaoAuctionGovDataRealtimeGetAPIResponseModel is 获取实时(今日)统计数据 成功返回结果
type TaobaoAuctionGovDataRealtimeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"auction_gov_data_realtime_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 今日拍卖统计数据
	RealTimeData *RealTimeData `json:"real_time_data,omitempty" xml:"real_time_data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAuctionGovDataRealtimeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RealTimeData = nil
}

var poolTaobaoAuctionGovDataRealtimeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAuctionGovDataRealtimeGetAPIResponse)
	},
}

// GetTaobaoAuctionGovDataRealtimeGetAPIResponse 从 sync.Pool 获取 TaobaoAuctionGovDataRealtimeGetAPIResponse
func GetTaobaoAuctionGovDataRealtimeGetAPIResponse() *TaobaoAuctionGovDataRealtimeGetAPIResponse {
	return poolTaobaoAuctionGovDataRealtimeGetAPIResponse.Get().(*TaobaoAuctionGovDataRealtimeGetAPIResponse)
}

// ReleaseTaobaoAuctionGovDataRealtimeGetAPIResponse 将 TaobaoAuctionGovDataRealtimeGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAuctionGovDataRealtimeGetAPIResponse(v *TaobaoAuctionGovDataRealtimeGetAPIResponse) {
	v.Reset()
	poolTaobaoAuctionGovDataRealtimeGetAPIResponse.Put(v)
}
