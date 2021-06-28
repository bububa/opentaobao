package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退票接口 APIResponse
alitrip.rail.trade.refund

退票接口
*/
type AlitripRailTradeRefundAPIResponse struct {
    model.CommonResponse
    AlitripRailTradeRefundResponse
}

type AlitripRailTradeRefundResponse struct {
    XMLName xml.Name `xml:"alitrip_rail_trade_refund_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回对象
    
    Result   *AlitripRailTradeRefundResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
