package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGet 根据mac查询设备的安全二维码
// alibaba.ailabs.tmallgenie.auth.device.withmac.qrcode.get
//
// 根据mac查询二维码详细信息
func AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGet(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIRequest, resp *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
