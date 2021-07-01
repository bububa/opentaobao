package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripSellerRefundGetAPIRequest
【机票代理商】退票申请单详情 API请求
taobao.alitrip.seller.refund.get

查询退票申请单详情 */
type TaobaoAlitripSellerRefundGetAPIRequest struct {
	model.Params
	// 申请单ID
	_applyId int64
}

// New
