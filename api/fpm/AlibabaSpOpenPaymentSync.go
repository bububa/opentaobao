package fpm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fpm"
)

// AlibabaSpOpenPaymentSync 付款单同步
// alibaba.sp.open.payment.sync
//
// 新康众弹外同步付款数据
func AlibabaSpOpenPaymentSync(clt *core.SDKClient, req *fpm.AlibabaSpOpenPaymentSyncAPIRequest, resp *fpm.AlibabaSpOpenPaymentSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
