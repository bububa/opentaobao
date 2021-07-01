package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemAdgroupRptdailylistAPIRequest
推广单元分日数据查询 API请求
taobao.feedflow.item.adgroup.rptdailylist

推广单元分日数据查询 */
type TaobaoFeedflowItemAdgroupRptdailylistAPIRequest struct {
	model.Params
	// 查询条件
	_rptQueryDTO *RptQueryDto
}

// New
