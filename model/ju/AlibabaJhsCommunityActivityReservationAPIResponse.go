package ju

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJhsCommunityActivityReservationAPIResponse 社群活动预约 API返回值
// alibaba.jhs.community.activity.reservation
//
// 社群活动预约
type AlibabaJhsCommunityActivityReservationAPIResponse struct {
	model.CommonResponse
	AlibabaJhsCommunityActivityReservationAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJhsCommunityActivityReservationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJhsCommunityActivityReservationAPIResponseModel).Reset()
}

// AlibabaJhsCommunityActivityReservationAPIResponseModel is 社群活动预约 成功返回结果
type AlibabaJhsCommunityActivityReservationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jhs_community_activity_reservation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// 是否预约成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaJhsCommunityActivityReservationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Error = ""
	m.Data = false
}

var poolAlibabaJhsCommunityActivityReservationAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJhsCommunityActivityReservationAPIResponse)
	},
}

// GetAlibabaJhsCommunityActivityReservationAPIResponse 从 sync.Pool 获取 AlibabaJhsCommunityActivityReservationAPIResponse
func GetAlibabaJhsCommunityActivityReservationAPIResponse() *AlibabaJhsCommunityActivityReservationAPIResponse {
	return poolAlibabaJhsCommunityActivityReservationAPIResponse.Get().(*AlibabaJhsCommunityActivityReservationAPIResponse)
}

// ReleaseAlibabaJhsCommunityActivityReservationAPIResponse 将 AlibabaJhsCommunityActivityReservationAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJhsCommunityActivityReservationAPIResponse(v *AlibabaJhsCommunityActivityReservationAPIResponse) {
	v.Reset()
	poolAlibabaJhsCommunityActivityReservationAPIResponse.Put(v)
}
