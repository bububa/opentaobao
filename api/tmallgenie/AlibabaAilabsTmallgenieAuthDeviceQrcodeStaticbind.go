package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbind 静态二维码绑定
// alibaba.ailabs.tmallgenie.auth.device.qrcode.staticbind
//
// 静态二维码绑定
func AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbind(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIRequest, session string) (*tmallgenie.AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse, error) {
	var resp tmallgenie.AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
