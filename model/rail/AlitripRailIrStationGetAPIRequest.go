package rail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripRailIrStationGetAPIRequest
国际火车票标准车站查询 API请求
alitrip.rail.ir.station.get

国际火车票提供给代理商用于查询标准车站信息，用于代理商对自己的车站与飞猪平台的车站做映射 */
type AlitripRailIrStationGetAPIRequest struct {
	model.Params
	// 商家id
	_agentId int64
	// 页数 从1开始
	_pageIndex int64
	// 每页条数
	_pageSize int64
}

// New
