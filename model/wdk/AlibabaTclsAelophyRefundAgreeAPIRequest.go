package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyRefundAgreeAPIRequest
saas 售后逆向 商户同意用户逆向申请 API请求
alibaba.tcls.aelophy.refund.agree

saas 售后逆向 商户同意用户逆向申请 */
type AlibabaTclsAelophyRefundAgreeAPIRequest struct {
	model.Params
	// 门店ID
	_storeId string
	// 外部订单ID
	_outOrderId string
	// 退款单ID
	_refundId string
	// 审核说明
	_auditMemo string
	// 外部子订单列表
	_subRefundList []Subrefundlist
}

// New
