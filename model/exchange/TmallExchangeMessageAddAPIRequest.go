package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallExchangeMessageAddAPIRequest
卖家创建换货留言 API请求
tmall.exchange.message.add

卖家创建换货留言 */
type TmallExchangeMessageAddAPIRequest struct {
	model.Params
	// 留言内容
	_content string
	// 换货单号ID
	_disputeId int64
	// 凭证图片列表
	_messagePics *model.File
	// 返回字段。目前支持id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type
	_fields []string
}

// New
