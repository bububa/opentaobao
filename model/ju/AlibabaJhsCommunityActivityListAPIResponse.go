package ju

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJhsCommunityActivityListAPIResponse 聚划算用增淘外社群服务活动列表 API返回值
// alibaba.jhs.community.activity.list
//
// 聚划算用增淘外社群服务活动列表
type AlibabaJhsCommunityActivityListAPIResponse struct {
	model.CommonResponse
	AlibabaJhsCommunityActivityListAPIResponseModel
}

// AlibabaJhsCommunityActivityListAPIResponseModel is 聚划算用增淘外社群服务活动列表 成功返回结果
type AlibabaJhsCommunityActivityListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jhs_community_activity_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Data []CommunityActivityDto `json:"data,omitempty" xml:"data>community_activity_dto,omitempty"`
	// 错误信息
	Error string `json:"error,omitempty" xml:"error,omitempty"`
}
