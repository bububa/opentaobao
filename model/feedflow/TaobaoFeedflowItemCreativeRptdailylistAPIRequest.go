package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCreativeRptdailylistAPIRequest
创意分日数据查询 API请求
taobao.feedflow.item.creative.rptdailylist

创意分日数据查询 */
type TaobaoFeedflowItemCreativeRptdailylistAPIRequest struct {
	model.Params
	// 查询条件
	_rptQueryDTO *RptQueryDto
}

// New
