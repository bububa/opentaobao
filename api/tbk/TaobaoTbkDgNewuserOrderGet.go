package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkdgnewuserorderget 淘宝客-推广者-新用户订单明细查询
// taobao.tbk.dg.newuser.order.get
//
// 拉新API
func Taobaotbkdgnewuserorderget(clt *core.SDKClient, req *tbk.TaobaotbkdgnewuserordergetAPIRequest, session string) (*tbk.TaobaotbkdgnewuserordergetAPIResponse, error) {
	var resp tbk.TaobaotbkdgnewuserordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
