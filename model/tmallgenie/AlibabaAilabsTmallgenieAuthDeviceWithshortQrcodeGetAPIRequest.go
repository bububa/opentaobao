package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIRequest
根据安全简码查询二维码详细信息 API请求
alibaba.ailabs.tmallgenie.auth.device.withshort.qrcode.get

根据安全简码查询二维码详细信息 */
type AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIRequest struct {
	model.Params
	// 产品ID
	_clientId string
	// 授权码
	_authCode string
}

// New
