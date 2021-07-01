package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallRefundModifyAPIRequest
修改OpenMall退款申请 API请求
taobao.openmall.refund.modify

修改OpenMall退款申请 */
type TaobaoOpenmallRefundModifyAPIRequest struct {
	model.Params
	// 退款类型，可选值refund（仅退款）、return_and_refund（退款退货）
	_refundType string
	// 退款金额，分
	_refundFee int64
	// 买家的退货描述
	_refundDesc string
	// 货品状态，可选值 BUYER_NOT_RECEIVED（买家未收到货）、BUYER_RECEIVED（买家已收到货）、UNSHIPPED（未发货）
	_goodsStatus string
	// 退款类别，可选值OTHER_REASON（其他，默认）、SEVEN_DAYS_WITHOUT_REASON（7天无理由，不退邮费）
	_refundReason string
	// 分销者联盟身份
	_distributor string
	// 退款单ID
	_refundId int64
}

// New
