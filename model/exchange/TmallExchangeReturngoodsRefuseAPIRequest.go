package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallExchangeReturngoodsRefuseAPIRequest
卖家拒绝确认收货 API请求
tmall.exchange.returngoods.refuse

卖家拒绝买家换货申请 */
type TmallExchangeReturngoodsRefuseAPIRequest struct {
	model.Params
	// 凭证图片
	_leaveMessagePics *model.File
	// 留言说明
	_leaveMessage string
	// 换货单号ID
	_disputeId int64
	// 拒绝原因ID
	_sellerRefuseReasonId int64
}

// New
