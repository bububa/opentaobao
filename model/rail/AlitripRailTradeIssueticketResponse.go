package rail

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
德铁出票成功接口 APIResponse
alitrip.rail.trade.issueticket

出票成功回调接口
*/
type AlitripRailTradeIssueticketAPIResponse struct {
    model.CommonResponse
    AlitripRailTradeIssueticketResponse
}

type AlitripRailTradeIssueticketResponse struct {
    XMLName xml.Name `xml:"alitrip_rail_trade_issueticket_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询结果集
    
    Result   *AlitripRailTradeIssueticketResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
