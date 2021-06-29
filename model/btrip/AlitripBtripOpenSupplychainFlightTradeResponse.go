package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】机票交易流水查询接口 APIResponse
alitrip.btrip.open.supplychain.flight.trade

【商旅】杭州市政府机票交易流水接口查询
*/
type AlitripBtripOpenSupplychainFlightTradeAPIResponse struct {
    model.CommonResponse
    AlitripBtripOpenSupplychainFlightTradeResponse
}

type AlitripBtripOpenSupplychainFlightTradeResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_open_supplychain_flight_trade_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *HisvResultSupport `json:"result,omitempty" xml:"result,omitempty"`

    
}
