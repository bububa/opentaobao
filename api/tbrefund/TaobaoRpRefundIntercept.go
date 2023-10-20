package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRpRefundIntercept 卖家发起拦截
// taobao.rp.refund.intercept
//
// 卖家发起拦截
func TaobaoRpRefundIntercept(clt *core.SDKClient, req *tbrefund.TaobaoRpRefundInterceptAPIRequest, resp *tbrefund.TaobaoRpRefundInterceptAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
