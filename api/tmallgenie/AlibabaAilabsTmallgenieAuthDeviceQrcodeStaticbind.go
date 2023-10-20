package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbind 静态二维码绑定
// alibaba.ailabs.tmallgenie.auth.device.qrcode.staticbind
//
// 静态二维码绑定
func AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbind(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIRequest, resp *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
