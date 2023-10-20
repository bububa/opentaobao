package hotelalliance

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelAllianceHidGetAPIResponse 获取联盟hid API返回值
// alitrip.hotel.alliance.hid.get
//
// 获取符合条件的菲住联盟hid，目前支持指定日期上线的菲住联盟hid查询
type AlitripHotelAllianceHidGetAPIResponse struct {
	model.CommonResponse
	AlitripHotelAllianceHidGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripHotelAllianceHidGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripHotelAllianceHidGetAPIResponseModel).Reset()
}

// AlitripHotelAllianceHidGetAPIResponseModel is 获取联盟hid 成功返回结果
type AlitripHotelAllianceHidGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_alliance_hid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	HmsTopResultSet *HmsTopResultSet `json:"hms_top_result_set,omitempty" xml:"hms_top_result_set,omitempty"`
}

// Reset 清空结构体
func (m *AlitripHotelAllianceHidGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.HmsTopResultSet = nil
}

var poolAlitripHotelAllianceHidGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripHotelAllianceHidGetAPIResponse)
	},
}

// GetAlitripHotelAllianceHidGetAPIResponse 从 sync.Pool 获取 AlitripHotelAllianceHidGetAPIResponse
func GetAlitripHotelAllianceHidGetAPIResponse() *AlitripHotelAllianceHidGetAPIResponse {
	return poolAlitripHotelAllianceHidGetAPIResponse.Get().(*AlitripHotelAllianceHidGetAPIResponse)
}

// ReleaseAlitripHotelAllianceHidGetAPIResponse 将 AlitripHotelAllianceHidGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripHotelAllianceHidGetAPIResponse(v *AlitripHotelAllianceHidGetAPIResponse) {
	v.Reset()
	poolAlitripHotelAllianceHidGetAPIResponse.Put(v)
}
