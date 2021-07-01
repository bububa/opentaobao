package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallChannelTradeOrderGetsAPIRequest
分页查询采购单 API请求
tmall.channel.trade.order.gets

分页查询采购单 */
type TmallChannelTradeOrderGetsAPIRequest struct {
	model.Params
	// 是否包含子单
	_isIncludeSubOrder bool
	// 是否包含主单
	_isIncludeMainOrder bool
	// 是否包含物流信息
	_isIncludeLogistics bool
	// 每页显示数量
	_pageSize int64
	// 查询第几页
	_pageNumber int64
	// 是否分页查询
	_needPagination bool
	// 主采购单号
	_mainPurchaseOrderNo int64
	// 分销商Nick
	_distributorNick string
	// 渠道编码
	_channel int64
	// 1-代销；2-经销
	_tradeType int64
	// 1. 待付款 2.已付款待发货 3.已发货待收货 4.交易完成 5.交易关闭
	_orderStatus int64
	// 创建时间从
	_createTimeStart string
	// 创建时间到
	_createTimeEnd string
}

// New
