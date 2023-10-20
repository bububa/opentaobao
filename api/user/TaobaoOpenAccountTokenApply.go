package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaoopenaccounttokenapply 申请免登Open Account Token
// taobao.open.account.token.apply
//
// 申请免登Open Account Token
func Taobaoopenaccounttokenapply(clt *core.SDKClient, req *user.TaobaoopenaccounttokenapplyAPIRequest, session string) (*user.TaobaoopenaccounttokenapplyAPIResponse, error) {
	var resp user.TaobaoopenaccounttokenapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
