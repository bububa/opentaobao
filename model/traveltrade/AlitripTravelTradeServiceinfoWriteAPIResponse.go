package traveltrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单服务信息写入接口 API返回值 
alitrip.travel.trade.serviceinfo.write

订单服务信息写入接口
*/
type AlitripTravelTradeServiceinfoWriteAPIResponse struct {
    model.CommonResponse
    AlitripTravelTradeServiceinfoWriteAPIResponseModel
}

// 订单服务信息写入接口 成功返回结果
type AlitripTravelTradeServiceinfoWriteAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_travel_trade_serviceinfo_write_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlitripTravelTradeServiceinfoWriteResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
