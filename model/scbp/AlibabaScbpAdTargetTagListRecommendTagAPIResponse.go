package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdTargetTagListRecommendTagAPIResponse 给计划推荐标签 API返回值
// alibaba.scbp.ad.target.tag.list.recommend.tag
//
// 给计划推荐标签
type AlibabaScbpAdTargetTagListRecommendTagAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdTargetTagListRecommendTagAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdTargetTagListRecommendTagAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdTargetTagListRecommendTagAPIResponseModel).Reset()
}

// AlibabaScbpAdTargetTagListRecommendTagAPIResponseModel is 给计划推荐标签 成功返回结果
type AlibabaScbpAdTargetTagListRecommendTagAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_target_tag_list_recommend_tag_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 定向标签推荐结果
	ResultList []TargetTagRecommendResultDto `json:"result_list,omitempty" xml:"result_list>target_tag_recommend_result_dto,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdTargetTagListRecommendTagAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
}

var poolAlibabaScbpAdTargetTagListRecommendTagAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdTargetTagListRecommendTagAPIResponse)
	},
}

// GetAlibabaScbpAdTargetTagListRecommendTagAPIResponse 从 sync.Pool 获取 AlibabaScbpAdTargetTagListRecommendTagAPIResponse
func GetAlibabaScbpAdTargetTagListRecommendTagAPIResponse() *AlibabaScbpAdTargetTagListRecommendTagAPIResponse {
	return poolAlibabaScbpAdTargetTagListRecommendTagAPIResponse.Get().(*AlibabaScbpAdTargetTagListRecommendTagAPIResponse)
}

// ReleaseAlibabaScbpAdTargetTagListRecommendTagAPIResponse 将 AlibabaScbpAdTargetTagListRecommendTagAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdTargetTagListRecommendTagAPIResponse(v *AlibabaScbpAdTargetTagListRecommendTagAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdTargetTagListRecommendTagAPIResponse.Put(v)
}
