package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripSellerRefundRefusereturnAPIRequest
【机票代理商】拒绝退票 API请求
taobao.alitrip.seller.refund.refusereturn

拒绝退票 */
type TaobaoAlitripSellerRefundRefusereturnAPIRequest struct {
	model.Params
	// 申请单ID
	_applyId int64
	// 拒绝理由
	_reason string
}

// New
