package brandhub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest
品牌号品牌特秀计划报表数据查询 API请求
taobao.brandhub.specialshow.rpt.campaign.get

获取品牌号品牌特秀广告campaign分日报表数据，只能查询近90天内的数据，包括展现量，点击量等 */
type TaobaoBrandhubSpecialshowRptCampaignGetAPIRequest struct {
	model.Params
	// 开始时间(最多可查询最近90天)
	_startDate string
	// 指定计划id
	_solutionId string
	// 截至时间(最晚到昨天)
	_endDate string
	// 当前页数
	_pageIndex string
	// 每页条数
	_pageSize string
}

// New
