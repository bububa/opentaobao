package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripRailTradeRefundAPIResponse
退票接口 API返回值
alitrip.rail.trade.refund

退票接口 */
type AlitripRailTradeRefundAPIResponse struct {
	model.CommonResponse
	AlitripRailTradeRefundAPIResponseModel
}

// AlitripRailTradeRefundAPIResponseModel is 退票接口 成功返回结果
type AlitripRailTradeRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_rail_trade_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回对象
	Result *AlitripRailTradeRefundResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
