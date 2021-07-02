package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoOpenAccountTokenValidate open account token验证
// taobao.open.account.token.validate
//
// open account token验证
func TaobaoOpenAccountTokenValidate(clt *core.SDKClient, req *user.TaobaoOpenAccountTokenValidateAPIRequest, session string) (*user.TaobaoOpenAccountTokenValidateAPIResponse, error) {
	var resp user.TaobaoOpenAccountTokenValidateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
