package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpcampaignfindlistAPIResponse 查询全量计划列表(不分页) API返回值
// taobao.universalbp.campaign.findlist
//
// 查询场景内的全量计划列表
type TaobaouniversalbpcampaignfindlistAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpcampaignfindlistAPIResponseModel
}

// TaobaouniversalbpcampaignfindlistAPIResponseModel is 查询全量计划列表(不分页) 成功返回结果
type TaobaouniversalbpcampaignfindlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_campaign_findlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpcampaignfindlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
