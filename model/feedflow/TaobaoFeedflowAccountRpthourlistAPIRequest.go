package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowAccountRpthourlistAPIRequest
超级推荐广告主分时报表查询 API请求
taobao.feedflow.account.rpthourlist

广告主分时报表查询，支持广告主查询最近90天内某一天的账户维度分时报表数据 */
type TaobaoFeedflowAccountRpthourlistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQuery *RptQueryDto
}

// New
