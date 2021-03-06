package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainFlightSearchAPIResponse 【商旅】机票订单查询 API返回值
// alitrip.btrip.supplychain.flight.search
//
// 【商旅】机票订单查询
type AlitripBtripSupplychainFlightSearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripSupplychainFlightSearchAPIResponseModel
}

// AlitripBtripSupplychainFlightSearchAPIResponseModel is 【商旅】机票订单查询 成功返回结果
type AlitripBtripSupplychainFlightSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_supplychain_flight_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
