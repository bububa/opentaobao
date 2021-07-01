package eleenterpriseordernew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest
设置订单支付 API请求
alibaba.ele.enterprise.ordernew.paymentstatus

设置订单支付成功 */
type AlibabaEleEnterpriseOrdernewPaymentstatusAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
	// 支付流水号
	_paySerialNumber string
}

// New
