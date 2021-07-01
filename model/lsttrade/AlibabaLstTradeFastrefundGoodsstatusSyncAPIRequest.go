package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest
卖家退款单商品状态同步 API请求
alibaba.lst.trade.fastrefund.goodsstatus.sync

卖家退款单商品状态同步 */
type AlibabaLstTradeFastrefundGoodsstatusSyncAPIRequest struct {
	model.Params
	// 主订单id
	_mainOrderId int64
	// 退款单id
	_refundId string
	// 未发货，枚举类型：UNSEND
	_status string
}

// New
