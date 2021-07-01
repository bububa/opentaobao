package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaInsightWordsareadataGetAPIRequest
获取关键词按地域进行细分的数据 API请求
taobao.simba.insight.wordsareadata.get

获取关键词按地域细分的详细数据，目前地域只能细化到省级别，返回的结果中包括市，是为了方便以后扩展，目前结果中市的值等于省。 */
type TaobaoSimbaInsightWordsareadataGetAPIRequest struct {
	model.Params
	// 关键词
	_bidword string
	// 开始时间，格式：yyyy-MM-dd
	_startDate string
	// 结束时间，格式：yyyy-MM-dd
	_endDate string
}

// New
