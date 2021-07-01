package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCampaignRpthourlistAPIResponse
超级推荐【商品推广】计划分时报表查询 API返回值
taobao.feedflow.item.campaign.rpthourlist

广告主推广计划分时数据查询，支持广告主查询最近90天内某一天的计划维度分时报表数据 */
type TaobaoFeedflowItemCampaignRpthourlistAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemCampaignRpthourlistAPIResponseModel
}

// TaobaoFeedflowItemCampaignRpthourlistAPIResponseModel is 超级推荐【商品推广】计划分时报表查询 成功返回结果
type TaobaoFeedflowItemCampaignRpthourlistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_campaign_rpthourlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TaobaoFeedflowItemCampaignRpthourlistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
