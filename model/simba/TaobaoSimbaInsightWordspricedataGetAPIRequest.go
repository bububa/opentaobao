package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaInsightWordspricedataGetAPIRequest
关键词按竞价区间的分布数据 API请求
taobao.simba.insight.wordspricedata.get

获取关键词按竞价区间进行细分的数据 */
type TaobaoSimbaInsightWordspricedataGetAPIRequest struct {
	model.Params
	// 关键词
	_bidword string
	// 开始时间，格式：yyyy-MM-dd
	_startDate string
	// 结束时间，格式：yyyy-MM-dd
	_endDate string
}

// New
