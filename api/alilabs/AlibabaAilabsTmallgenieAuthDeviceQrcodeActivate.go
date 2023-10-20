package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabsTmallgenieAuthDeviceQrcodeActivate 扫码激活设备
// alibaba.ailabs.tmallgenie.auth.device.qrcode.activate
//
// 三方带屏设备显示二维码（从天猫精灵云获取），使用三方APP扫码，将扫码到的安全code，通过TOP接口请求天猫精灵云，精灵云解析安全code的数据并激活对应的设备。
func AlibabaAilabsTmallgenieAuthDeviceQrcodeActivate(clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest, resp *alilabs.AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
