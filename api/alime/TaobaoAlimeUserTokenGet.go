package alime

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alime"
)

// TaobaoAlimeUserTokenGet 获取用户免登录令牌
// taobao.alime.user.token.get
//
// 根据第三账号信息获取用户的免登录令牌
func TaobaoAlimeUserTokenGet(ctx context.Context, clt *core.SDKClient, req *alime.TaobaoAlimeUserTokenGetAPIRequest, resp *alime.TaobaoAlimeUserTokenGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
