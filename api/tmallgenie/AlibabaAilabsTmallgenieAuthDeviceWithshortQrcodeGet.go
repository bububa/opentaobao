package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGet 根据安全简码查询二维码详细信息
// alibaba.ailabs.tmallgenie.auth.device.withshort.qrcode.get
//
// 根据安全简码查询二维码详细信息
func AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGet(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIRequest, session string) (*tmallgenie.AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse, error) {
	var resp tmallgenie.AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
