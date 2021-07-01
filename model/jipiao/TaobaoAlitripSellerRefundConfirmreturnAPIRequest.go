package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripSellerRefundConfirmreturnAPIRequest
【机票代理商】确认退票 API请求
taobao.alitrip.seller.refund.confirmreturn

确认退票 */
type TaobaoAlitripSellerRefundConfirmreturnAPIRequest struct {
	model.Params
	// 退票申请单
	_applyId int64
}

// New
