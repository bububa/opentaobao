package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieThirdTelecomAutoauth 电信iot自动授权
// alibaba.ailabs.tmallgenie.third.telecom.autoauth
//
// 电信iot自动授权
func AlibabaAilabsTmallgenieThirdTelecomAutoauth(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest, resp *tmallgenie.AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
