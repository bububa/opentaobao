package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsAligenieDeviceRegister 天猫精灵开放平台获取设备秘钥
// alibaba.ailabs.aligenie.device.register
//
// 向天猫精灵inside平台注册设备mac地址，并获取设备的唯一密钥
func AlibabaAilabsAligenieDeviceRegister(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsAligenieDeviceRegisterAPIRequest, session string) (*tmallgenie.AlibabaAilabsAligenieDeviceRegisterAPIResponse, error) {
	var resp tmallgenie.AlibabaAilabsAligenieDeviceRegisterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
