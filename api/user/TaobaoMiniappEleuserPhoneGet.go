package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoMiniappEleuserPhoneGet 获取饿了么用户信息
// taobao.miniapp.eleuser.phone.get
//
// 获取饿了么用户信息
func TaobaoMiniappEleuserPhoneGet(ctx context.Context, clt *core.SDKClient, req *user.TaobaoMiniappEleuserPhoneGetAPIRequest, resp *user.TaobaoMiniappEleuserPhoneGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
