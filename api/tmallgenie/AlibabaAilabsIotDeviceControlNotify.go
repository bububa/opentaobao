package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabsiotdevicecontrolnotify 天猫精灵IoT异步控制回调接口
// alibaba.ailabs.iot.device.control.notify
//
// 用于天猫精灵IoT云云接入控制结果的异步回调
func Alibabaailabsiotdevicecontrolnotify(clt *core.SDKClient, req *tmallgenie.AlibabaailabsiotdevicecontrolnotifyAPIRequest, session string) (*tmallgenie.AlibabaailabsiotdevicecontrolnotifyAPIResponse, error) {
	var resp tmallgenie.AlibabaailabsiotdevicecontrolnotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
