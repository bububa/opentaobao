package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorShake 摇一摇
// alibaba.interact.sensor.shake
//
// 摇一摇
func AlibabaInteractSensorShake(clt *core.SDKClient, req *interact.AlibabaInteractSensorShakeAPIRequest, session string) (*interact.AlibabaInteractSensorShakeAPIResponse, error) {
	var resp interact.AlibabaInteractSensorShakeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
