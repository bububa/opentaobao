package brandhub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBrandStartshopRptWordpackageGetAPIRequest
明星店铺品牌流量包报表数据查询 API请求
taobao.brand.startshop.rpt.wordpackage.get

获取明星店铺广告词包分日报表数据，只能查询近90天内的数据，包括展现量，点击量等 */
type TaobaoBrandStartshopRptWordpackageGetAPIRequest struct {
	model.Params
	// 开始日期
	_startDate string
	// 结束日期
	_endDate string
	// 转化周期
	_effect string
	// 流量类型
	_trafficType string
	// 每页显示条数(0,200]
	_pageSize string
	// 当前页数 ,从1开始
	_pageIndex string
}

// New
