package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCampaignCreateAPIResponse 创建计划 API返回值
// alibaba.scbp.ad.campaign.create
//
// 创建计划
type AlibabaScbpAdCampaignCreateAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdCampaignCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdCampaignCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdCampaignCreateAPIResponseModel).Reset()
}

// AlibabaScbpAdCampaignCreateAPIResponseModel is 创建计划 成功返回结果
type AlibabaScbpAdCampaignCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 计划id
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdCampaignCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolAlibabaScbpAdCampaignCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdCampaignCreateAPIResponse)
	},
}

// GetAlibabaScbpAdCampaignCreateAPIResponse 从 sync.Pool 获取 AlibabaScbpAdCampaignCreateAPIResponse
func GetAlibabaScbpAdCampaignCreateAPIResponse() *AlibabaScbpAdCampaignCreateAPIResponse {
	return poolAlibabaScbpAdCampaignCreateAPIResponse.Get().(*AlibabaScbpAdCampaignCreateAPIResponse)
}

// ReleaseAlibabaScbpAdCampaignCreateAPIResponse 将 AlibabaScbpAdCampaignCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdCampaignCreateAPIResponse(v *AlibabaScbpAdCampaignCreateAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdCampaignCreateAPIResponse.Put(v)
}
