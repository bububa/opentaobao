package alilabs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabUserOpenUidGet access token 获取精灵用户 id
// alibaba.ailab.user.open.uid.get
//
// access token 获取精灵用户 id
func AlibabaAilabUserOpenUidGet(ctx context.Context, clt *core.SDKClient, req *alilabs.AlibabaAilabUserOpenUidGetAPIRequest, resp *alilabs.AlibabaAilabUserOpenUidGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
