package brandhub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBrandStartshopRptAdgroupGetAPIRequest
明星店铺推广单元报表数据查询 API请求
taobao.brand.startshop.rpt.adgroup.get

获取明星店铺广告adgroup分日报表数据，只能查询近90天内的数据，包括展现量，点击量等 */
type TaobaoBrandStartshopRptAdgroupGetAPIRequest struct {
	model.Params
	// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
	_trafficType string
	// 转化周期默认15天,3,7,15
	_effect int64
	// 当前页数
	_pageIndex string
	// 每页条数
	_pageSize string
	// 开始时间(最多可查询最近90天)
	_startDate string
	// 截至时间(最晚到昨天)
	_endDate string
}

// New
