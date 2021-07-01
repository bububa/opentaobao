package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaInsightCatstopwordnewGetAPIRequest
获取类目下最热门的词 API请求
taobao.simba.insight.catstopwordnew.get

按照某个维度，查询某个类目下最热门的词，维度有点击，展现，花费，点击率等，具体可以按哪些字段进行排序，参考参数说明，比如选择了impression，则返回该类目下展现量最高那几个词。 */
type TaobaoSimbaInsightCatstopwordnewGetAPIRequest struct {
	model.Params
	// 类目id
	_catId string
	// 查询开始时间，格式必须为：yyyy-MM-dd
	_startDate string
	// 查询截止时间，格式只能是：yyyy-MM-dd
	_endDate string
	// 表示查询的维度，比如选择click，则查询该类目下点击量最大的词，可供选择的值有：impression, click, cost, ctr, cpc, coverage, transactiontotal, transactionshippingtotal, favtotal, roi
	_dimension string
	// 返回前多少条数据
	_pageSize int64
}

// New
