package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemAdzoneRptdailylistAPIRequest
资源包分日数据查询 API请求
taobao.feedflow.item.adzone.rptdailylist

资源包分日数据查询 */
type TaobaoFeedflowItemAdzoneRptdailylistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQueryDTO *RptQueryDto
}

// New
