package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoOrderConfirmPaidAPIRequest
确认收款 API请求
taobao.fenxiao.order.confirm.paid

供应商确认收款（非支付宝交易）。 */
type TaobaoFenxiaoOrderConfirmPaidAPIRequest struct {
	model.Params
	// 采购单编号。
	_purchaseOrderId int64
	// 确认支付信息（字数小于100）
	_confirmRemark string
}

// New
