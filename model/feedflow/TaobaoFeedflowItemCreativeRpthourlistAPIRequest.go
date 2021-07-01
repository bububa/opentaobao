package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCreativeRpthourlistAPIRequest
超级推荐【商品推广】创意分时报表查询 API请求
taobao.feedflow.item.creative.rpthourlist

创意分时数据查询，支持广告主查询最近90天内某一天的创意维度分时报表数据 */
type TaobaoFeedflowItemCreativeRpthourlistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQuery *RptQueryDto
}

// New
