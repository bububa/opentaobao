package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripSellerRefundFillfeeAPIRequest
机票代理商】回填手续费 API请求
taobao.alitrip.seller.refund.fillfee

回填手续费 */
type TaobaoAlitripSellerRefundFillfeeAPIRequest struct {
	model.Params
	// 申请单ID
	_applyId int64
	// 费对于关系，格式：{apply_fee_id:123,value:费用,金额单位分}
	_feePriceMap string
	// 改签费用，格式：{detail_id:123,value:费用,金额单位分}
	_modifyFee string
	// 票价信息，格式：{apply_fee_id：123,value:费用,金额单位分}
	_ticketPriceMap string
	// 升舱费用，格式：{detail_id:123,value:费用,金额单位分}
	_upgradeFee string
}

// New
