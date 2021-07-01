package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallChannelTradeApplyorderGetsAPIRequest
获取采购申请单列表 API请求
tmall.channel.trade.applyorder.gets

分页查询采购申请单列表 */
type TmallChannelTradeApplyorderGetsAPIRequest struct {
	model.Params
	// 交易类型
	_tradeType int64
	// 申请单号
	_channelPurchaseApplyOrderNo string
	// 每页显示数量
	_pageSize int64
	// 查询第几页
	_pageNumber int64
	// 是否需要分页
	_needPagination bool
	// 审核状态列表
	_auditStatusList []int64
	// 分销商nick
	_distributorNick string
	// 渠道
	_channel int64
}

// New
