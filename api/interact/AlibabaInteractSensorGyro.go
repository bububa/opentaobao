package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorGyro 陀螺仪
// alibaba.interact.sensor.gyro
//
// 客户端陀螺仪
func AlibabaInteractSensorGyro(clt *core.SDKClient, req *interact.AlibabaInteractSensorGyroAPIRequest, session string) (*interact.AlibabaInteractSensorGyroAPIResponse, error) {
	var resp interact.AlibabaInteractSensorGyroAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
