package brandhub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBrandStarshopRptTargetGetAPIRequest
明星店铺定向维度报表 API请求
taobao.brand.starshop.rpt.target.get

获取明星店铺定向维度分日报表数据，只能查询近90天内的数据，包括展现量，点击量等 */
type TaobaoBrandStarshopRptTargetGetAPIRequest struct {
	model.Params
	// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
	_trafficType string
	// 当前页数
	_pageIndex string
	// 每页条数
	_pageSize string
	// 转化周期,默认15, 3,7,15
	_effect string
	// 开始日期(最多查询1个月的数据)
	_startDate string
	// 截至日期(最晚到昨天)
	_endDate string
}

// New
