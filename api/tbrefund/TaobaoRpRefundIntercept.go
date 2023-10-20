package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// Taobaorprefundintercept 卖家发起拦截
// taobao.rp.refund.intercept
//
// 卖家发起拦截
func Taobaorprefundintercept(clt *core.SDKClient, req *tbrefund.TaobaorprefundinterceptAPIRequest, session string) (*tbrefund.TaobaorprefundinterceptAPIResponse, error) {
	var resp tbrefund.TaobaorprefundinterceptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
