package auction

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoauctiongovdatarealtimegetAPIResponse 获取实时(今日)统计数据 API返回值
// taobao.auction.gov.data.realtime.get
//
// 提供查询当日法院及下属法院的拍卖统计数据
type TaobaoauctiongovdatarealtimegetAPIResponse struct {
	model.CommonResponse
	TaobaoauctiongovdatarealtimegetAPIResponseModel
}

// TaobaoauctiongovdatarealtimegetAPIResponseModel is 获取实时(今日)统计数据 成功返回结果
type TaobaoauctiongovdatarealtimegetAPIResponseModel struct {
	XMLName xml.Name `xml:"auction_gov_data_realtime_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 今日拍卖统计数据
	RealTimeData *RealTimeData `json:"real_time_data,omitempty" xml:"real_time_data,omitempty"`
}
