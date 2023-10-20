package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse 删除屏蔽词 API返回值
// alibaba.scbp.ad.campaign.delete.forbidden.keyword
//
// 删除屏蔽词
type AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponseModel).Reset()
}

// AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponseModel is 删除屏蔽词 成功返回结果
type AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_delete_forbidden_keyword_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolAlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse)
	},
}

// GetAlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse 从 sync.Pool 获取 AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse
func GetAlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse() *AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse {
	return poolAlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse.Get().(*AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse)
}

// ReleaseAlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse 将 AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse(v *AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse.Put(v)
}
