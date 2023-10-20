package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Taobaowdkiotdeviceadminmqttdevicegetwithtoken 获取mqtt设备信息
// taobao.wdk.iot.deviceadmin.mqtt.device.getwithtoken
//
// 智能硬件设备动态注册和获取mqtt设备信息
func Taobaowdkiotdeviceadminmqttdevicegetwithtoken(clt *core.SDKClient, req *wdk.TaobaowdkiotdeviceadminmqttdevicegetwithtokenAPIRequest, session string) (*wdk.TaobaowdkiotdeviceadminmqttdevicegetwithtokenAPIResponse, error) {
	var resp wdk.TaobaowdkiotdeviceadminmqttdevicegetwithtokenAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
