package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallExchangeRefuseAPIRequest
卖家拒绝换货申请 API请求
tmall.exchange.refuse

卖家拒绝换货申请 */
type TmallExchangeRefuseAPIRequest struct {
	model.Params
	// 凭证图片
	_leaveMessagePics *model.File
	// 拒绝换货申请时的留言
	_leaveMessage string
	// 换货单号ID
	_disputeId int64
	// 换货原因对应ID
	_sellerRefuseReasonId int64
	// 返回字段。目前支持dispute_id, bizorder_id, modified, status
	_fields []string
}

// New
