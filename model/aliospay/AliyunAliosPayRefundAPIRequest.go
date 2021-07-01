package aliospay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunAliosPayRefundAPIRequest
退款接口 API请求
aliyun.alios.pay.refund

商户用来发起退款的接口 */
type AliyunAliosPayRefundAPIRequest struct {
	model.Params
	// 请求参数
	_refundRequest *RefundRequest
}

// New
