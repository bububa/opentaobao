package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsAligenieDeviceRegisterAPIRequest
天猫精灵开放平台获取设备秘钥 API请求
alibaba.ailabs.aligenie.device.register

向天猫精灵inside平台注册设备mac地址，并获取设备的唯一密钥 */
type AlibabaAilabsAligenieDeviceRegisterAPIRequest struct {
	model.Params
	// 设备id
	_deviceId int64
	// mac区段脚本
	_macSections string
}

// New
