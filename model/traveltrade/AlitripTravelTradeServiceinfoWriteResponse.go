package traveltrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单服务信息写入接口 APIResponse
alitrip.travel.trade.serviceinfo.write

订单服务信息写入接口
*/
type AlitripTravelTradeServiceinfoWriteAPIResponse struct {
    model.CommonResponse
    AlitripTravelTradeServiceinfoWriteResponse
}

type AlitripTravelTradeServiceinfoWriteResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_trade_serviceinfo_write_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlitripTravelTradeServiceinfoWriteResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
