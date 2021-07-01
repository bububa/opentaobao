package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaInsightCatsworddataGetAPIRequest
获取类目下关键词的数据 API请求
taobao.simba.insight.catsworddata.get

获取给定词在给定类目下的详细数据 */
type TaobaoSimbaInsightCatsworddataGetAPIRequest struct {
	model.Params
	// 类目id
	_catId string
	// 需要查询的关键词列表，最大长度100。
	_bidwordList []string
	// 开始时间，格式只能为：yyyy-MM-dd
	_startDate string
	// 结束时间，格式只能为：yyyy-MM-dd
	_endDate string
}

// New
