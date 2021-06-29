package rail

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
出票失败关单接口 API返回值 
alitrip.rail.trade.closeticket

出票成功回调接口
*/
type AlitripRailTradeCloseticketAPIResponse struct {
    model.CommonResponse
    AlitripRailTradeCloseticketResponse
}

// 出票失败关单接口 成功返回结果
type AlitripRailTradeCloseticketResponse struct {
    XMLName xml.Name `xml:"alitrip_rail_trade_closeticket_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 查询结果集
    Result   *AlitripRailTradeCloseticketResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
