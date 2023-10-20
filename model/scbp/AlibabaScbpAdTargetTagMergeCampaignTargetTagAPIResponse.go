package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse 标签增删改 API返回值
// alibaba.scbp.ad.target.tag.merge.campaign.target.tag
//
// 标签增删改
type AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponseModel).Reset()
}

// AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponseModel is 标签增删改 成功返回结果
type AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_target_tag_merge_campaign_target_tag_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolAlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse)
	},
}

// GetAlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse 从 sync.Pool 获取 AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse
func GetAlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse() *AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse {
	return poolAlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse.Get().(*AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse)
}

// ReleaseAlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse 将 AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse(v *AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse.Put(v)
}
