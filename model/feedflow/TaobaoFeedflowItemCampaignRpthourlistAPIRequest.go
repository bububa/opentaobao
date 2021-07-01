package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCampaignRpthourlistAPIRequest
超级推荐【商品推广】计划分时报表查询 API请求
taobao.feedflow.item.campaign.rpthourlist

广告主推广计划分时数据查询，支持广告主查询最近90天内某一天的计划维度分时报表数据 */
type TaobaoFeedflowItemCampaignRpthourlistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQuery *RptQueryDto
}

// New
