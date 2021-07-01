package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemAdgroupRpthourlistAPIRequest
超级推荐【商品推广】单元分时报表查询 API请求
taobao.feedflow.item.adgroup.rpthourlist

广告主推广组分时数据查询，支持广告主查询最近90天内某一天的单元维度分时报表数据 */
type TaobaoFeedflowItemAdgroupRpthourlistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQuery *RptQueryDto
}

// New
