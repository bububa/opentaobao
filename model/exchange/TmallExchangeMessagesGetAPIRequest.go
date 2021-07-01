package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallExchangeMessagesGetAPIRequest
查询换货订单留言列表 API请求
tmall.exchange.messages.get

查询换货订单留言列表 */
type TmallExchangeMessagesGetAPIRequest struct {
	model.Params
	// 留言创建角色。具体包括：卖家主账户(1)、卖家子账户(2)、小二(3)、买家(4)、系统(5)、系统超时(6)
	_operatorRoles []int64
	// 每页条数
	_pageSize int64
	// 换货单号ID
	_disputeId int64
	// 页码
	_pageNo int64
	// 返回的字段。具体包括：id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type
	_fields []string
}

// New
