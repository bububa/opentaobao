package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallChannelTradeDeliverorderGetsAPIRequest
查询发货单列表 API请求
tmall.channel.trade.deliverorder.gets

查询发货单列表 */
type TmallChannelTradeDeliverorderGetsAPIRequest struct {
	model.Params
	// 发货单单号
	_mainDeliverOrderNo int64
	// 发货单状态列表
	_orderStatusList []int64
	// 是否包括子发货单
	_isIncludeSubOrder bool
	// 每页显示数量
	_pageSize int64
	// 查询第几页
	_pageNumber int64
	// 是否分页查询
	_needPagination bool
	// 渠道
	_channel int64
}

// New
