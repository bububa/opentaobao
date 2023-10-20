package fpm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fpm"
)

// AlibabaSpOpenPaymentRepay 智付重新打款
// alibaba.sp.open.payment.repay
//
// 智付重新打款
func AlibabaSpOpenPaymentRepay(clt *core.SDKClient, req *fpm.AlibabaSpOpenPaymentRepayAPIRequest, session string) (*fpm.AlibabaSpOpenPaymentRepayAPIResponse, error) {
	var resp fpm.AlibabaSpOpenPaymentRepayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
