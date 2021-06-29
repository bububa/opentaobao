package traveltrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-订单发货接口 APIResponse
alitrip.travel.trade.deliver

航旅度假无需物流普通商品发货接口（不支持二次预约商品），只支持子订单级别发货
*/
type AlitripTravelTradeDeliverAPIResponse struct {
    model.CommonResponse
    AlitripTravelTradeDeliverResponse
}

type AlitripTravelTradeDeliverResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_trade_deliver_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 发货是否成功
    
    FirstResult   bool `json:"first_result,omitempty" xml:"first_result,omitempty"`

    
}
