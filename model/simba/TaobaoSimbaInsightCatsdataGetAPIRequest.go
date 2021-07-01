package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaInsightCatsdataGetAPIRequest
获取类目的大盘数据 API请求
taobao.simba.insight.catsdata.get

根据类目id获取类目的大盘数据，包括展现指数，点击指数，点击率，本次提供的insight相关的其它接口的都是这种情况。 */
type TaobaoSimbaInsightCatsdataGetAPIRequest struct {
	model.Params
	// 表示要查询的类目id
	_categoryIdList []string
	// 开始时间，格式：yyyy-MM-dd
	_startDate string
	// 查询截止时间，格式：yyyy-MM-dd
	_endDate string
}

// New
