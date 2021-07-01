package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取机票订单列表 API返回值 
alitrip.btrip.flight.order.search

第三方OA厂商获取机票订单列表
*/
type AlitripBtripFlightOrderSearchAPIResponse struct {
    model.CommonResponse
    AlitripBtripFlightOrderSearchAPIResponseModel
}

// 获取机票订单列表 成功返回结果
type AlitripBtripFlightOrderSearchAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_btrip_flight_order_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回数据
    Result   *BtriphomeResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}
