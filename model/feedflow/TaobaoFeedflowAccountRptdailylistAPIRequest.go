package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowAccountRptdailylistAPIRequest
获取广告主分日数据 API请求
taobao.feedflow.account.rptdailylist

获取广告主分日数据 */
type TaobaoFeedflowAccountRptdailylistAPIRequest struct {
	model.Params
	// 查询条件
	_rptQueryDTO *RptQueryDto
}

// New
