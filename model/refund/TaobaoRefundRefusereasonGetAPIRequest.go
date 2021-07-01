package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRefundRefusereasonGetAPIRequest
获取拒绝原因列表 API请求
taobao.refund.refusereason.get

获取商家拒绝原因列表 */
type TaobaoRefundRefusereasonGetAPIRequest struct {
	model.Params
	// 退款编号
	_refundId int64
	// 返回参数
	_fields string
	// 售中或售后
	_refundPhase string
}

// New
