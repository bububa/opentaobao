package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoOrderConfirmPaid 确认收款
// taobao.fenxiao.order.confirm.paid
//
// 供应商确认收款（非支付宝交易）。
func TaobaoFenxiaoOrderConfirmPaid(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoOrderConfirmPaidAPIRequest, resp *fenxiao.TaobaoFenxiaoOrderConfirmPaidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
