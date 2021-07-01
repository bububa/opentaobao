package brandhub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBrandStartshopRptAccountGetAPIRequest
明星店铺账户报表数据查询 API请求
taobao.brand.startshop.rpt.account.get

获取明星店铺广告主账户整体报表数据，只能查询近90天内的数据，包括展现量，点击量等 */
type TaobaoBrandStartshopRptAccountGetAPIRequest struct {
	model.Params
	// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
	_trafficType string
	// 默认15天
	_effect string
	// 开始时间(最多可查询最近90天)
	_endDate string
	// 截至时间(最晚到昨天)
	_startDate string
}

// New
