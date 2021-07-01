package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest
扫码激活设备 API请求
alibaba.ailabs.tmallgenie.auth.device.qrcode.activate

三方带屏设备显示二维码（从天猫精灵云获取），使用三方APP扫码，将扫码到的安全code，通过TOP接口请求天猫精灵云，精灵云解析安全code的数据并激活对应的设备。 */
type AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIRequest struct {
	model.Params
	// OAUTH authcode码
	_code string
	// 产品终端ID
	_clientId string
	// 淘宝授权ID
	_taobaoUserOpenid string
	// 扩展数据
	_extInfo string
}

// New
