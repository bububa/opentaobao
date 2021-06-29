package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商旅用车交易流水接口 APIResponse
alitrip.btrip.open.supplychain.vehicle.trade

商旅用车交易流水接口
*/
type AlitripBtripOpenSupplychainVehicleTradeAPIResponse struct {
    model.CommonResponse
    AlitripBtripOpenSupplychainVehicleTradeResponse
}

type AlitripBtripOpenSupplychainVehicleTradeResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_open_supplychain_vehicle_trade_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *HisvResultSupport `json:"result,omitempty" xml:"result,omitempty"`

    
}
