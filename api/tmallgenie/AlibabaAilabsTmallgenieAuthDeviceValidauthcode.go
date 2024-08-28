package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieAuthDeviceValidauthcode 根据authcode查询绑定结果
// alibaba.ailabs.tmallgenie.auth.device.validauthcode
//
// 根据authcode查询绑定结果
func AlibabaAilabsTmallgenieAuthDeviceValidauthcode(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest, resp *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
