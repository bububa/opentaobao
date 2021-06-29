package traveltrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-订单关闭接口（快速退款） APIResponse
alitrip.travel.trade.close

卖家关单（快速退款接口），不支持二次预约商品的订单
*/
type AlitripTravelTradeCloseAPIResponse struct {
    model.CommonResponse
    AlitripTravelTradeCloseResponse
}

type AlitripTravelTradeCloseResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_trade_close_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 交易关闭是否成功
    
    FirstResult   bool `json:"first_result,omitempty" xml:"first_result,omitempty"`

    
}
