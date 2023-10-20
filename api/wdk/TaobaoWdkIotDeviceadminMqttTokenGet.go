package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Taobaowdkiotdeviceadminmqtttokenget 获取MQTT访问令牌
// taobao.wdk.iot.deviceadmin.mqtt.token.get
//
// 智能硬件设备动态注册和获取mqtt设备信息
func Taobaowdkiotdeviceadminmqtttokenget(clt *core.SDKClient, req *wdk.TaobaowdkiotdeviceadminmqtttokengetAPIRequest, session string) (*wdk.TaobaowdkiotdeviceadminmqtttokengetAPIResponse, error) {
	var resp wdk.TaobaowdkiotdeviceadminmqtttokengetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
