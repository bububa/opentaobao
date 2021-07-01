package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaInsightWordssubdataGetAPIRequest
获取关键词按流量细分的数据 API请求
taobao.simba.insight.wordssubdata.get

获取关键词按流量进行细分的数据，返回结果中network表示流量的来源，意义如下：1->PC站内，2->PC站外,4->无线站内 5->无线站外 */
type TaobaoSimbaInsightWordssubdataGetAPIRequest struct {
	model.Params
	// 关键词列表
	_bidwordList []string
	// 开始时间
	_startDate string
	// 结束时间
	_endDate string
}

// New
