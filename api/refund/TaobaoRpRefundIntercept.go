package refund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/refund"
)

// TaobaoRpRefundIntercept 卖家发起拦截
// taobao.rp.refund.intercept
//
// 卖家发起拦截
func TaobaoRpRefundIntercept(clt *core.SDKClient, req *refund.TaobaoRpRefundInterceptAPIRequest, session string) (*refund.TaobaoRpRefundInterceptAPIResponse, error) {
	var resp refund.TaobaoRpRefundInterceptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
