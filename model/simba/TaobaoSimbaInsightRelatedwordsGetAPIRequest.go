package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaInsightRelatedwordsGetAPIRequest
获取词的相关词 API请求
taobao.simba.insight.relatedwords.get

获取给定词的若干相关词，返回结果中越相关的权重越大，排在越前面，根据number参数对返回结果进行截断。 */
type TaobaoSimbaInsightRelatedwordsGetAPIRequest struct {
	model.Params
	// 要查询的词列表
	_bidwordList []string
	// 表示返回数据的条数
	_number int64
}

// New
