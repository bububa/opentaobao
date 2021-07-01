package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripSellerRefundmoneyConfirmAPIRequest
【机票代理商订单】确认退款 API请求
taobao.alitrip.seller.refundmoney.confirm

代理人确认退票申请单的退款 */
type TaobaoAlitripSellerRefundmoneyConfirmAPIRequest struct {
	model.Params
	// 申请单id
	_applyId int64
}

// New
