package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallRefundCloseAPIRequest
关闭OpenMall退款单 API请求
taobao.openmall.refund.close

关闭OpenMall退款单 */
type TaobaoOpenmallRefundCloseAPIRequest struct {
	model.Params
	// 渠道
	_distributor string
	// 退款ID
	_refundId int64
}

// New
