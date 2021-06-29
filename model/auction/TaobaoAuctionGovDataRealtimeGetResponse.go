package auction

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取实时(今日)统计数据 APIResponse
taobao.auction.gov.data.realtime.get

提供查询当日法院及下属法院的拍卖统计数据
*/
type TaobaoAuctionGovDataRealtimeGetAPIResponse struct {
    model.CommonResponse
    TaobaoAuctionGovDataRealtimeGetResponse
}

type TaobaoAuctionGovDataRealtimeGetResponse struct {
    XMLName xml.Name `xml:"auction_gov_data_realtime_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 今日拍卖统计数据
    
    RealTimeData   *RealTimeData `json:"real_time_data,omitempty" xml:"real_time_data,omitempty"`

    
}
