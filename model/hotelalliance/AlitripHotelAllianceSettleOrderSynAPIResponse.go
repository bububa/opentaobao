package hotelalliance

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelAllianceSettleOrderSynAPIResponse 菲住联盟分账成功订单同步 API返回值
// alitrip.hotel.alliance.settle.order.syn
//
// 用于菲住联盟分账成功订单同步
type AlitripHotelAllianceSettleOrderSynAPIResponse struct {
	model.CommonResponse
	AlitripHotelAllianceSettleOrderSynAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripHotelAllianceSettleOrderSynAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripHotelAllianceSettleOrderSynAPIResponseModel).Reset()
}

// AlitripHotelAllianceSettleOrderSynAPIResponseModel is 菲住联盟分账成功订单同步 成功返回结果
type AlitripHotelAllianceSettleOrderSynAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_alliance_settle_order_syn_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的结果
	HmsTopResultSet *HmsTopResultSet `json:"hms_top_result_set,omitempty" xml:"hms_top_result_set,omitempty"`
}

// Reset 清空结构体
func (m *AlitripHotelAllianceSettleOrderSynAPIResponseModel) Reset() {
	m.RequestId = ""
	m.HmsTopResultSet = nil
}

var poolAlitripHotelAllianceSettleOrderSynAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripHotelAllianceSettleOrderSynAPIResponse)
	},
}

// GetAlitripHotelAllianceSettleOrderSynAPIResponse 从 sync.Pool 获取 AlitripHotelAllianceSettleOrderSynAPIResponse
func GetAlitripHotelAllianceSettleOrderSynAPIResponse() *AlitripHotelAllianceSettleOrderSynAPIResponse {
	return poolAlitripHotelAllianceSettleOrderSynAPIResponse.Get().(*AlitripHotelAllianceSettleOrderSynAPIResponse)
}

// ReleaseAlitripHotelAllianceSettleOrderSynAPIResponse 将 AlitripHotelAllianceSettleOrderSynAPIResponse 保存到 sync.Pool
func ReleaseAlitripHotelAllianceSettleOrderSynAPIResponse(v *AlitripHotelAllianceSettleOrderSynAPIResponse) {
	v.Reset()
	poolAlitripHotelAllianceSettleOrderSynAPIResponse.Put(v)
}
