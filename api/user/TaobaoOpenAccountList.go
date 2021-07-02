package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoOpenAccountList OpenAccount账号信息查询
// taobao.open.account.list
//
// OpenAccount账号信息查询
func TaobaoOpenAccountList(clt *core.SDKClient, req *user.TaobaoOpenAccountListAPIRequest, session string) (*user.TaobaoOpenAccountListAPIResponse, error) {
	var resp user.TaobaoOpenAccountListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
