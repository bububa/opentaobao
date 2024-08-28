package mtopopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// TaobaoOauthCodeCreate 淘宝OauthCode颁发
// taobao.oauth.code.create
//
// 手淘无线开放的oauthCode颁发接口
func TaobaoOauthCodeCreate(ctx context.Context, clt *core.SDKClient, req *mtopopen.TaobaoOauthCodeCreateAPIRequest, resp *mtopopen.TaobaoOauthCodeCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
