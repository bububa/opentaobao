package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCampaignDeleteAPIResponse 删除计划 API返回值
// alibaba.scbp.ad.campaign.delete
//
// 删除计划
type AlibabaScbpAdCampaignDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdCampaignDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdCampaignDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdCampaignDeleteAPIResponseModel).Reset()
}

// AlibabaScbpAdCampaignDeleteAPIResponseModel is 删除计划 成功返回结果
type AlibabaScbpAdCampaignDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除成功条数
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdCampaignDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolAlibabaScbpAdCampaignDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdCampaignDeleteAPIResponse)
	},
}

// GetAlibabaScbpAdCampaignDeleteAPIResponse 从 sync.Pool 获取 AlibabaScbpAdCampaignDeleteAPIResponse
func GetAlibabaScbpAdCampaignDeleteAPIResponse() *AlibabaScbpAdCampaignDeleteAPIResponse {
	return poolAlibabaScbpAdCampaignDeleteAPIResponse.Get().(*AlibabaScbpAdCampaignDeleteAPIResponse)
}

// ReleaseAlibabaScbpAdCampaignDeleteAPIResponse 将 AlibabaScbpAdCampaignDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdCampaignDeleteAPIResponse(v *AlibabaScbpAdCampaignDeleteAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdCampaignDeleteAPIResponse.Put(v)
}
