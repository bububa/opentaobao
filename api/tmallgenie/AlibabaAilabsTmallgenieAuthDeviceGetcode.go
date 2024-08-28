package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieAuthDeviceGetcode 获取authcode
// alibaba.ailabs.tmallgenie.auth.device.getcode
//
// 获取绑定的authcode
func AlibabaAilabsTmallgenieAuthDeviceGetcode(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIRequest, resp *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
