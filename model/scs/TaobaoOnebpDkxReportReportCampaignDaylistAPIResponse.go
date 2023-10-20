package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxreportreportcampaigndaylistAPIResponse 获取计划分日报表 API返回值
// taobao.onebp.dkx.report.report.campaign.daylist
//
// 获取计划分日报表，场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe
type TaobaoonebpdkxreportreportcampaigndaylistAPIResponse struct {
	model.CommonResponse
	TaobaoonebpdkxreportreportcampaigndaylistAPIResponseModel
}

// TaobaoonebpdkxreportreportcampaigndaylistAPIResponseModel is 获取计划分日报表 成功返回结果
type TaobaoonebpdkxreportreportcampaigndaylistAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_report_report_campaign_daylist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoonebpdkxreportreportcampaigndaylistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
