package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoRefundGetAPIRequest
查询采购单退款信息 API请求
taobao.fenxiao.refund.get

分销商或供应商可以查询某子单的退款信息，以及下游订单的退款信息 */
type TaobaoFenxiaoRefundGetAPIRequest struct {
	model.Params
	// 要查询的退款子单的id
	_subOrderId int64
	// 是否查询下游买家的退款信息
	_querySellerRefund bool
}

// New
