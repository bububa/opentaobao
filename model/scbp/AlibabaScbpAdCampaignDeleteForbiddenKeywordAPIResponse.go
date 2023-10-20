package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadcampaigndeleteforbiddenkeywordAPIResponse 删除屏蔽词 API返回值
// alibaba.scbp.ad.campaign.delete.forbidden.keyword
//
// 删除屏蔽词
type AlibabascbpadcampaigndeleteforbiddenkeywordAPIResponse struct {
	model.CommonResponse
	AlibabascbpadcampaigndeleteforbiddenkeywordAPIResponseModel
}

// AlibabascbpadcampaigndeleteforbiddenkeywordAPIResponseModel is 删除屏蔽词 成功返回结果
type AlibabascbpadcampaigndeleteforbiddenkeywordAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_delete_forbidden_keyword_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}
