package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGet 根据安全简码查询二维码详细信息
// alibaba.ailabs.tmallgenie.auth.device.withshort.qrcode.get
//
// 根据安全简码查询二维码详细信息
func AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGet(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIRequest, resp *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
