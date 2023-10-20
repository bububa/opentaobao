package ju

import (
	"encoding/xml"

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
