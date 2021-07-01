package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallChannelTradeRefundorderGetsAPIRequest
供应商查询退款单 API请求
tmall.channel.trade.refundorder.gets

供应商分页查询退款单 */
type TmallChannelTradeRefundorderGetsAPIRequest struct {
	model.Params
	// 退款单号
	_refundId int64
	// 采购单号
	_mainChannelOrderNo string
	// 每页数据条数
	_pageSize int64
	// 页码
	_pageNumber int64
}

// New
