package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRpReturngoodsRefillAPIRequest
卖家回填物流信息 API请求
taobao.rp.returngoods.refill

卖家收到货物回填物流信息，如果买家已经回填物流信息，则接口报错，目前仅支持天猫订单。 */
type TaobaoRpReturngoodsRefillAPIRequest struct {
	model.Params
	// 退款单编号
	_refundId int64
	// 退款阶段，可选值：售中：onsale，售后：aftersale
	_refundPhase string
	// 物流公司运单号
	_logisticsWaybillNo string
	// 物流公司编号
	_logisticsCompanyCode string
}

// New
