package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorBlow 吹气
// alibaba.interact.sensor.blow
//
// 客户端吹气
func AlibabaInteractSensorBlow(clt *core.SDKClient, req *interact.AlibabaInteractSensorBlowAPIRequest, session string) (*interact.AlibabaInteractSensorBlowAPIResponse, error) {
	var resp interact.AlibabaInteractSensorBlowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
