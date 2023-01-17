package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRpRefundIntercept 卖家发起拦截
// taobao.rp.refund.intercept
//
// 卖家发起拦截
func TaobaoRpRefundIntercept(clt *core.SDKClient, req *tbrefund.TaobaoRpRefundInterceptAPIRequest, session string) (*tbrefund.TaobaoRpRefundInterceptAPIResponse, error) {
	var resp tbrefund.TaobaoRpRefundInterceptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
