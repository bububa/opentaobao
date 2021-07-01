package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

/* AlibabaAilabsAligenieIotDeviceControlResult
设备控制结果
alibaba.ailabs.aligenie.iot.device.control.result

智能IOT解决外部厂商在云云模式在用户通过天猫精灵下发设备指令过程中，厂商指令完成，回调结果通知 */
func AlibabaAilabsAligenieIotDeviceControlResult(clt *core.SDKClient, req *iot.AlibabaAilabsAligenieIotDeviceControlResultAPIRequest, session string) (*iot.AlibabaAilabsAligenieIotDeviceControlResultAPIResponse, error) {
	var resp iot.AlibabaAilabsAligenieIotDeviceControlResultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
