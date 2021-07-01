package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallRefundCreateAPIRequest
创建OpenMall退款单 API请求
taobao.openmall.refund.create

创建OpenMall退款单
如存在未完结的退款单，则返回该退款单ID */
type TaobaoOpenmallRefundCreateAPIRequest struct {
	model.Params
	// 分销者联盟身份
	_distributor string
	// 货品状态，可选值 BUYER_NOT_RECEIVED（买家未收到货）、BUYER_RECEIVED（买家已收到货）、UNSHIPPED（未发货）
	_goodsStatus string
	// 买家的退货描述
	_refundDesc string
	// 退款金额，分
	_refundFee int64
	// 退款类别，可选值OTHER_REASON（其他）、SEVEN_DAYS_WITHOUT_REASON（7天无理由，不退邮费）
	_refundReason string
	// 退款类型，可选值refund（仅退款）、return_and_refund（退款退货）
	_refundType string
	// 订单号
	_tid int64
}

// New
