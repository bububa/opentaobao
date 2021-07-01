package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServiceSettleadjustmentSearchAPIRequest
服务商15分钟获取数据 API请求
tmall.service.settleadjustment.search

天猫服务平台，按修改时间，时间间隔在15中内（包含15分钟），获取调整单数据 */
type TmallServiceSettleadjustmentSearchAPIRequest struct {
	model.Params
	// 结束时间
	_endTime string
	// 开始时间
	_startTime string
}

// New
