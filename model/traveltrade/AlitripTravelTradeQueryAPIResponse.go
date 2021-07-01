package traveltrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-订单详情查询接口 API返回值 
alitrip.travel.trade.query

飞猪度假订单详情查询接口
*/
type AlitripTravelTradeQueryAPIResponse struct {
    model.CommonResponse
    AlitripTravelTradeQueryAPIResponseModel
}

// 飞猪度假-订单详情查询接口 成功返回结果
type AlitripTravelTradeQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_travel_trade_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 交易主订单详情
    FirstResult   *TopTripOrderResult `json:"first_result,omitempty" xml:"first_result,omitempty"`
}
