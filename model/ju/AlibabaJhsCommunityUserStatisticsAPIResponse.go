package ju

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJhsCommunityUserStatisticsAPIResponse 聚划算社群用户行为上报 API返回值
// alibaba.jhs.community.user.statistics
//
// 聚划算社群用户行为上报
type AlibabaJhsCommunityUserStatisticsAPIResponse struct {
	model.CommonResponse
	AlibabaJhsCommunityUserStatisticsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJhsCommunityUserStatisticsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJhsCommunityUserStatisticsAPIResponseModel).Reset()
}

// AlibabaJhsCommunityUserStatisticsAPIResponseModel is 聚划算社群用户行为上报 成功返回结果
type AlibabaJhsCommunityUserStatisticsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jhs_community_user_statistics_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// 上报数据是否成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaJhsCommunityUserStatisticsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Error = ""
	m.Data = false
}

var poolAlibabaJhsCommunityUserStatisticsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJhsCommunityUserStatisticsAPIResponse)
	},
}

// GetAlibabaJhsCommunityUserStatisticsAPIResponse 从 sync.Pool 获取 AlibabaJhsCommunityUserStatisticsAPIResponse
func GetAlibabaJhsCommunityUserStatisticsAPIResponse() *AlibabaJhsCommunityUserStatisticsAPIResponse {
	return poolAlibabaJhsCommunityUserStatisticsAPIResponse.Get().(*AlibabaJhsCommunityUserStatisticsAPIResponse)
}

// ReleaseAlibabaJhsCommunityUserStatisticsAPIResponse 将 AlibabaJhsCommunityUserStatisticsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJhsCommunityUserStatisticsAPIResponse(v *AlibabaJhsCommunityUserStatisticsAPIResponse) {
	v.Reset()
	poolAlibabaJhsCommunityUserStatisticsAPIResponse.Put(v)
}
