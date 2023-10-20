package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGet 根据安全简码查询二维码详细信息
// alibaba.ailabs.tmallgenie.auth.device.withshort.qrcode.get
//
// 根据安全简码查询二维码详细信息
func AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGet(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIRequest, resp *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
