package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripIeAgentOrderSearchAPIRequest
【国际机票】订单列表查询 API请求
taobao.alitrip.ie.agent.order.search

根据指定条件查询国际机票订单列表 */
type TaobaoAlitripIeAgentOrderSearchAPIRequest struct {
	model.Params
	// 代理商ID
	_agentId int64
	// 订单起始日期
	_beginTime string
	// 当前页码
	_currentPage int64
	// 订单结束日期
	_endTime string
	// 订单状态（只能传入一个状态，不支持多个一起传）
	_orderStatus string
	// 分页大小
	_pageSize int64
	// 0:自有运价;3:公布运价;9:大卖家API;11私有库存
	_fareSource int64
	// 供应渠道/资源码
	_resourceCode string
	// officeNo
	_officeNo string
}

// New
