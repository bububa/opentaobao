package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorGyro 陀螺仪
// alibaba.interact.sensor.gyro
//
// 客户端陀螺仪
func AlibabaInteractSensorGyro(clt *core.SDKClient, req *interact.AlibabaInteractSensorGyroAPIRequest, resp *interact.AlibabaInteractSensorGyroAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
