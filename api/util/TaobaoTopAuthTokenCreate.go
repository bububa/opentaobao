package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoTopAuthTokenCreate 获取Access Token
// taobao.top.auth.token.create
//
// 用户通过code换获取access_token，https only
func TaobaoTopAuthTokenCreate(ctx context.Context, clt *core.SDKClient, req *util.TaobaoTopAuthTokenCreateAPIRequest, resp *util.TaobaoTopAuthTokenCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
