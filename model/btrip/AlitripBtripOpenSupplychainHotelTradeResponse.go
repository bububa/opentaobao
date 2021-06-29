package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】酒店交易查询流水接口 APIResponse
alitrip.btrip.open.supplychain.hotel.trade

【商旅】酒店交易查询流水接口——杭州市政府
*/
type AlitripBtripOpenSupplychainHotelTradeAPIResponse struct {
    model.CommonResponse
    AlitripBtripOpenSupplychainHotelTradeResponse
}

type AlitripBtripOpenSupplychainHotelTradeResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_open_supplychain_hotel_trade_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *HisvResultSupport `json:"result,omitempty" xml:"result,omitempty"`

    
}
