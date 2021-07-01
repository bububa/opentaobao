package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCampaignRptdailylistAPIRequest
推广计划分日数据查询 API请求
taobao.feedflow.item.campaign.rptdailylist

推广计划分日数据查询 */
type TaobaoFeedflowItemCampaignRptdailylistAPIRequest struct {
	model.Params
	// 查询条件
	_rptQueryDTO *RptQueryDto
}

// New
