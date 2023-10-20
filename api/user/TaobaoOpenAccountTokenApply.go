package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoOpenAccountTokenApply 申请免登Open Account Token
// taobao.open.account.token.apply
//
// 申请免登Open Account Token
func TaobaoOpenAccountTokenApply(clt *core.SDKClient, req *user.TaobaoOpenAccountTokenApplyAPIRequest, session string) (*user.TaobaoOpenAccountTokenApplyAPIResponse, error) {
	var resp user.TaobaoOpenAccountTokenApplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
