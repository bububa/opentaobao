package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoOpenAccountDelete OpenAccount删除数据
// taobao.open.account.delete
//
// OpenAccount删除数据
func TaobaoOpenAccountDelete(clt *core.SDKClient, req *user.TaobaoOpenAccountDeleteAPIRequest, session string) (*user.TaobaoOpenAccountDeleteAPIResponse, error) {
	var resp user.TaobaoOpenAccountDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
