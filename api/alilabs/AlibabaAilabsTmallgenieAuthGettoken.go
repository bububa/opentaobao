package alilabs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabsTmallgenieAuthGettoken 设备授权
// alibaba.ailabs.tmallgenie.auth.gettoken
//
// 获取设备授权码
func AlibabaAilabsTmallgenieAuthGettoken(ctx context.Context, clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthGettokenAPIRequest, resp *alilabs.AlibabaAilabsTmallgenieAuthGettokenAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
