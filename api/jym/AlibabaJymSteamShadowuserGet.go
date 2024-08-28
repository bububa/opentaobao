package jym

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymSteamShadowuserGet 获取影子标识
// alibaba.jym.steam.shadowuser.get
//
// 交易猫Steam类目获取影子ID
func AlibabaJymSteamShadowuserGet(ctx context.Context, clt *core.SDKClient, req *jym.AlibabaJymSteamShadowuserGetAPIRequest, resp *jym.AlibabaJymSteamShadowuserGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
