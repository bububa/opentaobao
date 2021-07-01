package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemAdzoneRpthourlistAPIRequest
超级推荐【商品推广】资源位分时报表查询 API请求
taobao.feedflow.item.adzone.rpthourlist

广告主资源包分时数据查询，支持广告主查询最近90天内某一天的资源包维度分时报表数据 */
type TaobaoFeedflowItemAdzoneRpthourlistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQuery *RptQueryDto
}

// New
