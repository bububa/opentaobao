package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
退票接口 APIResponse
alitrip.rail.trade.refund

退票接口
*/
type AlitripRailTradeRefundAPIResponse struct {
    model.CommonResponse
    Response *AlitripRailTradeRefundResponse `json:"alitrip_rail_trade_refund_response,omitempty"`
}

type AlitripRailTradeRefundResponse struct {

    // 接口返回对象
    Result   *AlitripRailTradeRefundResultSet `json:"result,omitempty"`

}
