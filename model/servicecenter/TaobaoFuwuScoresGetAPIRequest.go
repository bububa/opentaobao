package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFuwuScoresGetAPIRequest
服务平台评价查询接口 API请求
taobao.fuwu.scores.get

根据日期、查询appkey对应服务评价，每次调用只能查询某一天服务评价信息，可设置分页查询，页大小最大为100，非实时接口，延迟时间为30分钟 */
type TaobaoFuwuScoresGetAPIRequest struct {
	model.Params
	// 当前页
	_currentPage int64
	// 每页获取条数。默认值40，最小值1，最大值100。
	_pageSize int64
	// 评价日期，查询某一天的评价
	_date string
}

// New
