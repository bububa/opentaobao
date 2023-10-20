package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaoopenaccountlist OpenAccount账号信息查询
// taobao.open.account.list
//
// OpenAccount账号信息查询
func Taobaoopenaccountlist(clt *core.SDKClient, req *user.TaobaoopenaccountlistAPIRequest, session string) (*user.TaobaoopenaccountlistAPIResponse, error) {
	var resp user.TaobaoopenaccountlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
