package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaoopenaccounttokenvalidate open account token验证
// taobao.open.account.token.validate
//
// open account token验证
func Taobaoopenaccounttokenvalidate(clt *core.SDKClient, req *user.TaobaoopenaccounttokenvalidateAPIRequest, session string) (*user.TaobaoopenaccounttokenvalidateAPIResponse, error) {
	var resp user.TaobaoopenaccounttokenvalidateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
