package brandhub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest
品牌号品牌特秀单元报表数据查询 API请求
taobao.brandhub.specialshow.rpt.adgroup.get

获取品牌号品牌特秀广告adgroup分日报表数据，只能查询近90天内的数据，包括展现量，点击量等 */
type TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest struct {
	model.Params
	// 开始时间(最多可查询最近90天)
	_startDate string
	// 截至时间(最晚到昨天)
	_endDate string
	// 指定计划id
	_solutionId string
	// 指定任务id
	_taskId string
	// 当前页数
	_pageIndex string
	// 可选页数
	_pageSize string
}

// New
