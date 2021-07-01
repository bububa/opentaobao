package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCrowdRptdailylistAPIRequest
定向分日数据查询 API请求
taobao.feedflow.item.crowd.rptdailylist

定向分日数据查询 */
type TaobaoFeedflowItemCrowdRptdailylistAPIRequest struct {
	model.Params
	// 查询条件
	_rptQueryDTO *RptQueryDto
}

// New
