package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadtargettagmergecampaigntargettagAPIResponse 标签增删改 API返回值
// alibaba.scbp.ad.target.tag.merge.campaign.target.tag
//
// 标签增删改
type AlibabascbpadtargettagmergecampaigntargettagAPIResponse struct {
	model.CommonResponse
	AlibabascbpadtargettagmergecampaigntargettagAPIResponseModel
}

// AlibabascbpadtargettagmergecampaigntargettagAPIResponseModel is 标签增删改 成功返回结果
type AlibabascbpadtargettagmergecampaigntargettagAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_target_tag_merge_campaign_target_tag_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}
