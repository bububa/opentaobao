package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaInsightCatsforecastnewGetAPIRequest
获取词的相关类目预测数据 API请求
taobao.simba.insight.catsforecastnew.get

根据给定的词，预测这些词的相关类目 */
type TaobaoSimbaInsightCatsforecastnewGetAPIRequest struct {
	model.Params
	// 需要查询的词列表
	_bidwordList []string
}

// New
