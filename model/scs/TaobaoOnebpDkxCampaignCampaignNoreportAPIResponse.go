package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxcampaigncampaignnoreportAPIResponse 获取场景计划的非报表数据 API返回值
// taobao.onebp.dkx.campaign.campaign.noreport
//
// 获取场景计划的非报表数据。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoonebpdkxcampaigncampaignnoreportAPIResponse struct {
	model.CommonResponse
	TaobaoonebpdkxcampaigncampaignnoreportAPIResponseModel
}

// TaobaoonebpdkxcampaigncampaignnoreportAPIResponseModel is 获取场景计划的非报表数据 成功返回结果
type TaobaoonebpdkxcampaigncampaignnoreportAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_campaign_campaign_noreport_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoonebpdkxcampaigncampaignnoreportResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
