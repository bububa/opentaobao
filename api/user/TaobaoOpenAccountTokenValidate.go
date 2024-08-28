package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoOpenAccountTokenValidate open account token验证
// taobao.open.account.token.validate
//
// open account token验证
func TaobaoOpenAccountTokenValidate(ctx context.Context, clt *core.SDKClient, req *user.TaobaoOpenAccountTokenValidateAPIRequest, resp *user.TaobaoOpenAccountTokenValidateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
