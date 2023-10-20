package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCampaignUpdateAPIResponse 修改计划 API返回值
// alibaba.scbp.ad.campaign.update
//
// 修改计划
type AlibabaScbpAdCampaignUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdCampaignUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdCampaignUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdCampaignUpdateAPIResponseModel).Reset()
}

// AlibabaScbpAdCampaignUpdateAPIResponseModel is 修改计划 成功返回结果
type AlibabaScbpAdCampaignUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改成功数
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdCampaignUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolAlibabaScbpAdCampaignUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdCampaignUpdateAPIResponse)
	},
}

// GetAlibabaScbpAdCampaignUpdateAPIResponse 从 sync.Pool 获取 AlibabaScbpAdCampaignUpdateAPIResponse
func GetAlibabaScbpAdCampaignUpdateAPIResponse() *AlibabaScbpAdCampaignUpdateAPIResponse {
	return poolAlibabaScbpAdCampaignUpdateAPIResponse.Get().(*AlibabaScbpAdCampaignUpdateAPIResponse)
}

// ReleaseAlibabaScbpAdCampaignUpdateAPIResponse 将 AlibabaScbpAdCampaignUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdCampaignUpdateAPIResponse(v *AlibabaScbpAdCampaignUpdateAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdCampaignUpdateAPIResponse.Put(v)
}
