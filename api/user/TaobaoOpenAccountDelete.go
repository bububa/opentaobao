package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoOpenAccountDelete OpenAccount删除数据
// taobao.open.account.delete
//
// OpenAccount删除数据
func TaobaoOpenAccountDelete(clt *core.SDKClient, req *user.TaobaoOpenAccountDeleteAPIRequest, resp *user.TaobaoOpenAccountDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
