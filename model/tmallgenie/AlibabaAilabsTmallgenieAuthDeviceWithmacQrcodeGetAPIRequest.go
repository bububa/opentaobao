package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIRequest
根据mac查询设备的安全二维码 API请求
alibaba.ailabs.tmallgenie.auth.device.withmac.qrcode.get

根据mac查询二维码详细信息 */
type AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIRequest struct {
	model.Params
	// 产品ID
	_clientId string
	// 设备mac地址
	_mac string
}

// New
