package alilabs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabsTmallgenieAuthGetcode 获取token
// alibaba.ailabs.tmallgenie.auth.getcode
//
// 获取天猫精灵authCode
func AlibabaAilabsTmallgenieAuthGetcode(ctx context.Context, clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthGetcodeAPIRequest, resp *alilabs.AlibabaAilabsTmallgenieAuthGetcodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
