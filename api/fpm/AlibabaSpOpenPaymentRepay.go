package fpm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fpm"
)

// AlibabaSpOpenPaymentRepay 智付重新打款
// alibaba.sp.open.payment.repay
//
// 智付重新打款
func AlibabaSpOpenPaymentRepay(clt *core.SDKClient, req *fpm.AlibabaSpOpenPaymentRepayAPIRequest, resp *fpm.AlibabaSpOpenPaymentRepayAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
