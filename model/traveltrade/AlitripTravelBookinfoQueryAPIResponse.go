package traveltrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelBookinfoQueryAPIResponse 飞猪度假-订单二次预约查询接口 API返回值
// alitrip.travel.bookinfo.query
//
// 飞猪度假订单二次预约详情查询接口
type AlitripTravelBookinfoQueryAPIResponse struct {
	model.CommonResponse
	AlitripTravelBookinfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelBookinfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelBookinfoQueryAPIResponseModel).Reset()
}

// AlitripTravelBookinfoQueryAPIResponseModel is 飞猪度假-订单二次预约查询接口 成功返回结果
type AlitripTravelBookinfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_bookinfo_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 交易预定结果对象
	FirstResult *TopTripBookInfoResult `json:"first_result,omitempty" xml:"first_result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelBookinfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FirstResult = nil
}

var poolAlitripTravelBookinfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelBookinfoQueryAPIResponse)
	},
}

// GetAlitripTravelBookinfoQueryAPIResponse 从 sync.Pool 获取 AlitripTravelBookinfoQueryAPIResponse
func GetAlitripTravelBookinfoQueryAPIResponse() *AlitripTravelBookinfoQueryAPIResponse {
	return poolAlitripTravelBookinfoQueryAPIResponse.Get().(*AlitripTravelBookinfoQueryAPIResponse)
}

// ReleaseAlitripTravelBookinfoQueryAPIResponse 将 AlitripTravelBookinfoQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelBookinfoQueryAPIResponse(v *AlitripTravelBookinfoQueryAPIResponse) {
	v.Reset()
	poolAlitripTravelBookinfoQueryAPIResponse.Put(v)
}
