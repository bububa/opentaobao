package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaInsightWordsdataGetAPIRequest
获取关键词的大盘数据 API请求
taobao.simba.insight.wordsdata.get

获取关键词的详细数据 */
type TaobaoSimbaInsightWordsdataGetAPIRequest struct {
	model.Params
	// 关键词列表，最多可传100个。
	_bidwordList []string
	// 开始时间
	_startDate string
	// 结束时间
	_endDate string
}

// New
