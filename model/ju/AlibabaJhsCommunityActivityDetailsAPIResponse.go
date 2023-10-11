package ju

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJhsCommunityActivityDetailsAPIResponse 社群活动详情 API返回值
// alibaba.jhs.community.activity.details
//
// 社群活动详情
type AlibabaJhsCommunityActivityDetailsAPIResponse struct {
	model.CommonResponse
	AlibabaJhsCommunityActivityDetailsAPIResponseModel
}

// AlibabaJhsCommunityActivityDetailsAPIResponseModel is 社群活动详情 成功返回结果
type AlibabaJhsCommunityActivityDetailsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jhs_community_activity_details_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// 数据对象
	Data *AlibabaJhsCommunityActivityDetailsData `json:"data,omitempty" xml:"data,omitempty"`
}
