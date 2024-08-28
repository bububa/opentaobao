package alilabs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabUserTokenGet 三方账号获取 token
// alibaba.ailab.user.token.get
//
// inside 设备的三方 app，通过 extId、schema 生成 token
func AlibabaAilabUserTokenGet(ctx context.Context, clt *core.SDKClient, req *alilabs.AlibabaAilabUserTokenGetAPIRequest, resp *alilabs.AlibabaAilabUserTokenGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
