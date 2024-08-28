package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbind 静态二维码绑定
// alibaba.ailabs.tmallgenie.auth.device.qrcode.staticbind
//
// 静态二维码绑定
func AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbind(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIRequest, resp *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
