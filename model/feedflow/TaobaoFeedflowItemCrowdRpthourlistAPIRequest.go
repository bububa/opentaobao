package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCrowdRpthourlistAPIRequest
超级推荐【商品推广】定向分时报表查询 API请求
taobao.feedflow.item.crowd.rpthourlist

广告主定向分时数据查询，支持广告主查询最近90天内某一天的定向维度分时报表数据 */
type TaobaoFeedflowItemCrowdRpthourlistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQuery *RptQueryDto
}

// New
