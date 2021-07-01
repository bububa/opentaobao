package rail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripRailIrServiceGetAPIRequest
国际火车票仓位坐席查询 API请求
alitrip.rail.ir.service.get

国际火车票标准仓位坐席查询 */
type AlitripRailIrServiceGetAPIRequest struct {
	model.Params
	// 6代表境外火车票
	_bizType int64
	// 代理商id
	_agentId int64
}

// New
