package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCustomersSidGet 查看功能权限
// taobao.simba.customers.sid.get
//
// 查询用户是否拥有某个功能权限
func TaobaoSimbaCustomersSidGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaCustomersSidGetAPIRequest, resp *simba.TaobaoSimbaCustomersSidGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
