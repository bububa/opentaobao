package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaHourReportAccountGetAPIRequest
账户级别小时报表获取 API请求
taobao.simba.hour.report.account.get

获取账户小时实时报表数据 */
type TaobaoSimbaHourReportAccountGetAPIRequest struct {
	model.Params
	// 昵称
	_nick string
	// 时间
	_theDate string
	// 当前小时
	_hour string
}

// New
