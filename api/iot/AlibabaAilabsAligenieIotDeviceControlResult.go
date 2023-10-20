package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Alibabaailabsaligenieiotdevicecontrolresult 设备控制结果
// alibaba.ailabs.aligenie.iot.device.control.result
//
// 智能IOT解决外部厂商在云云模式在用户通过天猫精灵下发设备指令过程中，厂商指令完成，回调结果通知
func Alibabaailabsaligenieiotdevicecontrolresult(clt *core.SDKClient, req *iot.AlibabaailabsaligenieiotdevicecontrolresultAPIRequest, session string) (*iot.AlibabaailabsaligenieiotdevicecontrolresultAPIResponse, error) {
	var resp iot.AlibabaailabsaligenieiotdevicecontrolresultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
