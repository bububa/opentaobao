package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaRptCampaigneffectGetAPIRequest
推广计划效果报表数据对象 API请求
taobao.simba.rpt.campaigneffect.get

推广计划效果报表数据对象 */
type TaobaoSimbaRptCampaigneffectGetAPIRequest struct {
	model.Params
	// 权限校验参数
	_subwayToken string
	// 昵称
	_nick string
	// 开始时间，格式yyyy-mm-dd
	_startTime string
	// 结束时间，格式yyyy-mm-dd
	_endTime string
	// 推广计划id
	_campaignId int64
	// 报表类型（搜索：SEARCH,类目出价：CAT,定向投放：NOSEARCH 全部：ALL）
	_searchType string
	// 数据来源（站内：1，站外：2）可多选以逗号分隔，默认值为：1,2
	_source string
	// 页码
	_pageNo int64
	// 每页大小
	_pageSize int64
}

// New
