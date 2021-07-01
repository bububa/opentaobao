package rail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripRailIrCarrierGetAPIRequest
国际火车票铁路承运公司查询 API请求
alitrip.rail.ir.carrier.get

国际火车票提供给代理商用于查询标准铁路承运公司carrier信息，用于代理商自己的carrier与飞猪平台的carrier做映射 */
type AlitripRailIrCarrierGetAPIRequest struct {
	model.Params
	// 商家id
	_agentId int64
}

// New
