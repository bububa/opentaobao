package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripVehicleOrderSearchAPIResponse 用车订单查询接口 API返回值
// alitrip.btrip.vehicle.order.search
//
// 企业获取商旅用车订单数据
type AlitripBtripVehicleOrderSearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripVehicleOrderSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripVehicleOrderSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripVehicleOrderSearchAPIResponseModel).Reset()
}

// AlitripBtripVehicleOrderSearchAPIResponseModel is 用车订单查询接口 成功返回结果
type AlitripBtripVehicleOrderSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_vehicle_order_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripVehicleOrderSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripVehicleOrderSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripVehicleOrderSearchAPIResponse)
	},
}

// GetAlitripBtripVehicleOrderSearchAPIResponse 从 sync.Pool 获取 AlitripBtripVehicleOrderSearchAPIResponse
func GetAlitripBtripVehicleOrderSearchAPIResponse() *AlitripBtripVehicleOrderSearchAPIResponse {
	return poolAlitripBtripVehicleOrderSearchAPIResponse.Get().(*AlitripBtripVehicleOrderSearchAPIResponse)
}

// ReleaseAlitripBtripVehicleOrderSearchAPIResponse 将 AlitripBtripVehicleOrderSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripVehicleOrderSearchAPIResponse(v *AlitripBtripVehicleOrderSearchAPIResponse) {
	v.Reset()
	poolAlitripBtripVehicleOrderSearchAPIResponse.Put(v)
}
