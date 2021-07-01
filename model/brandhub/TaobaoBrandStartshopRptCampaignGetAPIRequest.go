package brandhub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBrandStartshopRptCampaignGetAPIRequest
明星店铺推广计划报表数据查询 API请求
taobao.brand.startshop.rpt.campaign.get

获取明星店铺广告campaign分日报表数据，只能查询近90天内的数据，包括展现量，点击量等 */
type TaobaoBrandStartshopRptCampaignGetAPIRequest struct {
	model.Params
	// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
	_traffictype string
	// 查询开始时间(最多查询90天数据)
	_startdate string
	// 查询截至时间(最晚查询到昨天)
	_enddate string
	// 转化周期,默认15天,可选 3,7,15
	_effect int64
}

// New
